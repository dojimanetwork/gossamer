// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// VotingRuleResult A future returned by a `VotingRule` to restrict a given vote, if any restriction is necessary.
type VotingRuleResult struct{}

// VotingRule A trait for custom voting rules in GRANDPA.
type VotingRule struct {
	// this is a trait
}

// Restrict the given `current_target` vote, returning the block hash and
// number of the block to vote on, and `None` in case the vote should not
// be restricted. `base` is the block that we're basing our votes on in
// order to pick our target (e.g. last round estimate), and `best_target`
// is the initial best vote target before any vote rules were applied. When
// applying multiple `VotingRule`s both `base` and `best_target` should
// remain unchanged.
//
// The contract of this interface requires that when restricting a vote, the
// returned value **must** be an ancestor of the given `current_target`,
// this also means that a variant must be maintained throughout the
// execution of voting rules wherein `current_target <= best_target`.
func (s *VotingRule) restrict_vote() {}

// A custom voting rule that guarantees that our vote is always behind the best
// block by at least N blocks, unless the base number is < N blocks behind the
// best, in which case it votes for the base.
//
// In the best case our vote is exactly N blocks
// behind the best block, but if there is a scenario where either
// >34% of validators run without this rule or the fork-choice rule
// can prioritize shorter chains over longer ones, the vote may be
// closer to the best block than N.
type BeforeBestBlockBy struct{}

func (s *BeforeBestBlockBy) restrict_vote() {}

// A custom voting rule that limits votes towards 3/4 of the unfinalized chain,
// using the given `base` and `best_target` to figure where the 3/4 target
// should fall.
type ThreeQuartersOfTheUnfinalizedChain struct{}

func (s *ThreeQuartersOfTheUnfinalizedChain) restrict_vote() {}

// walk backwards until we find the target block
func find_target() {
	// TODO easy one
}

type VotingRules struct {
	//rules: Arc<Vec<Box<dyn VotingRule<Block, B>>>>,
}

func (s *VotingRules) restrict_vote() {}

// VotingRulesBuilder A builder of a composite voting rule that applies a set of rules to
// progressively restrict the vote.
type VotingRulesBuilder struct{}

func (s *VotingRulesBuilder) Default() {}

// Return a new voting rule builder using the given backend.
func (s *VotingRulesBuilder) new() {}

// Add a new voting rule to the builder.
func (s *VotingRulesBuilder) add() {}

// Add all given voting rules to the builder.
func (s *VotingRulesBuilder) add_all() {}

// Return a new `VotingRule` that applies all of the previously added
// voting rules in-order.
func (s *VotingRulesBuilder) build() {}

func (s *VotingRule) restrict_vote2() {}
