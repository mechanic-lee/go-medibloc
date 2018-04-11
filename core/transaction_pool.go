package core

import (
	"sync"

	"sort"

	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/common/hashheap"
	"github.com/medibloc/go-medibloc/util/logging"
	"github.com/sirupsen/logrus"
)

// TransactionPool is a pool of all received transactions from network.
type TransactionPool struct {
	mu sync.RWMutex

	size int

	candidates *hashheap.HashedHeap
	buckets    *hashheap.HashedHeap
	all        map[common.Hash]*Transaction
}

// NewTransactionPool returns TransactionPool.
func NewTransactionPool(size int) *TransactionPool {
	return &TransactionPool{
		size:       size,
		candidates: hashheap.NewHashedHeap(),
		buckets:    hashheap.NewHashedHeap(),
		all:        make(map[common.Hash]*Transaction),
	}
}

// Get returns transaction by tx hash.
func (pool *TransactionPool) Get(hash common.Hash) *Transaction {
	pool.mu.RLock()
	defer pool.mu.RUnlock()

	return pool.all[hash]
}

// Del deletes transaction.
func (pool *TransactionPool) Del(tx *Transaction) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	pool.del(tx)
}

// Push pushes transaction to the pool.
func (pool *TransactionPool) Push(tx *Transaction) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if err := pool.push(tx); err != nil {
		return err
	}

	pool.evict()

	return nil
}

// Pop pop transaction from the pool.
func (pool *TransactionPool) Pop() *Transaction {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	cmpTx := pool.candidates.Peek()
	if cmpTx == nil {
		return nil
	}
	tx := cmpTx.(*comparable).Transaction

	pool.del(tx)

	return tx
}

func (pool *TransactionPool) push(tx *Transaction) error {
	// push to all
	if _, ok := pool.all[tx.Hash()]; ok {
		return ErrDuplicatedTransaction
	}
	pool.all[tx.Hash()] = tx

	from := tx.From().Str()

	// push to bucket
	var bkt *bucket
	v := pool.buckets.Get(from)
	if v != nil {
		pool.buckets.Del(from)
		bkt = v.(*bucket)
	} else {
		bkt = newBucket()
	}
	bkt.push(tx)
	pool.buckets.Set(from, bkt)

	// replace candidate
	candidate := bkt.peekFirst()
	pool.candidates.Del(from)
	pool.candidates.Set(from, &comparable{candidate})

	return nil
}

func (pool *TransactionPool) del(tx *Transaction) {
	// Remove from all
	if _, ok := pool.all[tx.Hash()]; !ok {
		return
	}
	delete(pool.all, tx.Hash())

	from := tx.From().Str()

	// Remove from bucket
	v := pool.buckets.Get(from)
	if v == nil {
		logging.WithFields(logrus.Fields{
			"tx": tx,
		}).Error("Unable to find the bucket containing the transaction.")
		return
	}
	bkt := v.(*bucket)
	bkt.del(tx)

	// Remove from candidates
	pool.candidates.Del(from)

	// Remove bucket if empty
	if bkt.isEmpty() {
		return
	}
	pool.buckets.Set(from, bkt)

	// Replace candidate
	candidate := bkt.peekFirst()
	pool.candidates.Set(from, &comparable{candidate})
}

func (pool *TransactionPool) evict() {
	if len(pool.all) <= pool.size {
		return
	}

	// Peek longest bucket
	v := pool.buckets.Peek()
	if v == nil {
		return
	}
	bkt := v.(*bucket)

	tx := bkt.peekLast()
	if tx == nil {
		return
	}

	pool.del(tx)
}

// comparable assigns a sequence to the transaction.
type comparable struct{ *Transaction }

func (tx *comparable) Less(o interface{}) bool { return tx.nonce < o.(*comparable).nonce }

// transactions is sortable slice of comparable transactions.
type transactions []*comparable

func (t transactions) Len() int { return len(t) }

func (t transactions) Less(i, j int) bool { return t[i].Less(t[j]) }

func (t transactions) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// bucket is a set of transactions for each account.
type bucket struct {
	txs transactions
}

func newBucket() *bucket {
	return &bucket{
		txs: make(transactions, 0),
	}
}

func (b *bucket) Less(o interface{}) bool {
	return len(b.txs) > len(o.(*bucket).txs)
}

func (b *bucket) isEmpty() bool {
	return len(b.txs) == 0
}

func (b *bucket) push(tx *Transaction) {
	cmpTx := &comparable{tx}
	b.txs = append(b.txs, cmpTx)
	sort.Sort(b.txs)
}

func (b *bucket) peekFirst() *Transaction {
	if len(b.txs) == 0 {
		return nil
	}
	return b.txs[0].Transaction
}

func (b *bucket) peekLast() *Transaction {
	if len(b.txs) == 0 {
		return nil
	}
	return b.txs[len(b.txs)-1].Transaction
}

func (b *bucket) del(tx *Transaction) {
	for i, tt := range b.txs {
		if tt.Transaction.Hash() == tx.Hash() {
			b.txs = append(b.txs[:i], b.txs[i+1:]...)
			return
		}
	}
}