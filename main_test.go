// main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-1, -1)
    if result != -2 {
        t.Errorf("Add(-1, -1) = %d; want -2", result)
    }
}