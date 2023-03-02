//go:build !schema

package main

func schema() {
	panic("You must generate json schema with tag 'schema'")
}
