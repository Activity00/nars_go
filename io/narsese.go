package io

import (
	"nars_go/language"
	"strings"
)

/**
 * Utility methods for working and reacting to Narsese input.
 * This will eventually be integrated with NarseseParser for systematic
 * parsing and prediction of input.
 */

type Narsese struct {
}

func NewNarsese() *Narsese {
	return &Narsese{}
}

// ParseTerm /*  ---------- react String into term ---------- */
/**
 * Top-level method that react a Term in general, which may recursively call itself.
 * <p>
 * There are 5 valid cases: 1. (Op, A1, ..., An) is a CompoundTerm if Op is
 * a built-in getOperator 2. {A1, ..., An} is an SetExt; 3. [A1, ..., An] is an
 * SetInt; 4. &lt;T1 Re T2&gt; is a Statement (including higher-order Statement);
 * 5. otherwise it is a simple term.
 *
 * @param s the String to be parsed
 * @return the Term generated from the String
 *
 *
 */
func (n Narsese) ParseTerm(s string) (*language.Term, error) {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return nil, nil
	}

	return &language.Term{}, nil
}
