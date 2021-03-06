// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package core

import (
	"sort"

	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/util"
)

type candidate struct {
	address    common.Address
	votesPower *util.Uint128
	candidacy  bool
}

type candidates []*candidate

type votesCache struct {
	candidates candidates
}

func newVotesCache() *votesCache {
	return &votesCache{
		candidates: candidates{},
	}
}

func (cs candidates) Len() int {
	return len(cs)
}

func (cs candidates) Less(i, j int) bool {
	return cs[i].votesPower.Cmp(cs[j].votesPower) < 0
}

func (cs candidates) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (vc *votesCache) Clone() *votesCache {
	var cds candidates
	for _, cand := range vc.candidates {
		cds = append(cds, &candidate{
			address:    cand.address,
			votesPower: cand.votesPower.DeepCopy(),
			candidacy:  cand.candidacy,
		})
	}
	return &votesCache{
		candidates: cds,
	}
}

func (vc *votesCache) GetCandidate(address common.Address) (int, *candidate, error) {
	var idx int
	for idx = 0; idx < vc.candidates.Len(); idx++ {
		if vc.candidates[idx].address == address {
			break
		}
	}
	if idx == vc.candidates.Len() {
		return -1, nil, ErrCandidateNotFound
	}
	return idx, vc.candidates[idx], nil
}

func (vc *votesCache) AddCandidate(address common.Address, votesPower *util.Uint128) {
	_, c, err := vc.GetCandidate(address)
	if err == nil {
		c.candidacy = true
		return
	}
	c = &candidate{
		address:    address,
		votesPower: votesPower,
		candidacy:  true,
	}
	vc.candidates = append(vc.candidates, c)
	sort.Sort(vc.candidates)
}

func (vc *votesCache) RemoveCandidate(address common.Address) error {
	idx, _, err := vc.GetCandidate(address)
	if err != nil {
		return nil
	}
	vc.candidates = append(vc.candidates[:idx], vc.candidates[idx+1:]...)
	return nil
}

func (vc *votesCache) AddVotesPower(address common.Address, amount *util.Uint128) error {
	_, c, err := vc.GetCandidate(address)
	if err != nil && err != ErrCandidateNotFound {
		return err
	}
	if err == ErrCandidateNotFound {
		vc.AddCandidate(address, amount)
		return nil
	}
	c.votesPower, err = c.votesPower.Add(amount)
	if err != nil {
		return err
	}
	sort.Sort(vc.candidates)
	return nil
}

func (vc *votesCache) SubtractVotesPower(address common.Address, amount *util.Uint128) error {
	_, c, err := vc.GetCandidate(address)
	if err != nil {
		return err
	}
	if c.votesPower.Cmp(amount) < 0 {
		return ErrVotesPowerGetsMinus
	}
	c.votesPower, err = c.votesPower.Sub(amount)
	if err != nil {
		return err
	}
	sort.Sort(vc.candidates)
	return nil
}

func (vc *votesCache) SetCandidacy(address common.Address, candidacy bool) error {
	_, c, err := vc.GetCandidate(address)
	if err != nil {
		return err
	}
	c.candidacy = candidacy
	return nil
}
