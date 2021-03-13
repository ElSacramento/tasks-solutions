package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTrie(t *testing.T) {
	{
		words := []string{"foobar", "football", "fake", "kate", "a"}
		myTrie := CreateTrie(words)
		require.Equal(t, 3, len(myTrie.children))
	}
	{
		words := []string{"foobar"}
		require.Equal(t, 1, len(CreateTrie(words).children))
	}
	{
		words := []string{"foobar", "football", "fake"}
		require.Equal(t, 1, len(CreateTrie(words).children))
	}
}

func TestCamelMatch(t *testing.T) {
	{
		queries := []string{"FooBar", "FootBall", "Folks", "FooBarTest", "Pool"}
		pattern := "FB"
		expected := []bool{true, true, false, false, false}
		require.Equal(t, expected, camelMatch(queries, pattern))
	}
	{
		queries := []string{"FooBar", "FootBall", "Folks", "FooBarTest", "Pool", "FitnessBar", "FoB"}
		pattern := "FoB"
		expected := []bool{true, true, false, false, false, false, true}
		require.Equal(t, expected, camelMatch(queries, pattern))
	}
	{
		queries := []string{"FooBar", "FootBall", "Folks", "FooBarTest", "Pool"}
		pattern := "FoBaT"
		expected := []bool{false, false, false, true, false}
		require.Equal(t, expected, camelMatch(queries, pattern))
	}
	{
		queries := []string{"FooBar", "FootBall", "Folks", "FooBarTest", "Pool"}
		pattern := "M"
		expected := []bool{false, false, false, false, false}
		require.Equal(t, expected, camelMatch(queries, pattern))
	}
	{
		queries := []string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}
		pattern := "CooP"
		expected := []bool{false, false, true}
		require.Equal(t, expected, camelMatch(queries, pattern))
	}
}
