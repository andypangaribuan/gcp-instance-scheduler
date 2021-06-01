package test

import "testing"


/* ============================================
	Created by andy pangaribuan on 2021/05/19
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func logFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("\n\n:: ERROR\n%+v\n", err)
	}
}