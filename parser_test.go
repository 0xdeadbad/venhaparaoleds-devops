package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCParser(t *testing.T) {

	t.Run("Test cParser Str", func(t *testing.T) {
		c := newCParser("hello,world")
		vt1, v1 := c.Next()
		assert.Equal(t, Str, vt1, "type should been Str")
		assert.Equal(t, "hello", v1, "type value should been \"hello\"")

		vt2, v2 := c.Next()
		assert.Equal(t, Str, vt2, "type should been Str")
		assert.Equal(t, "world", v2, "type value should been \"hello\"")

		eof, _ := c.Next()
		assert.Equal(t, Eof, eof, "last Next() should return Eof")
	})

	t.Run("Test cParser StrList", func(t *testing.T) {
		c := newCParser("[hello,world]")
		vt1, v1 := c.Next()
		assert.Equal(t, ListStr, vt1, "type should been ListStr")
		assert.Equal(t, []string{"hello", "world"}, v1, "type value should been \"[hello,world]\"")

		eof, _ := c.Next()
		assert.Equal(t, Eof, eof, "last Next() should return Eof")
	})

	t.Run("Test cParser Str && StrList", func(t *testing.T) {
		c := newCParser("hello,[hello,world]")
		vt1, v1 := c.Next()
		assert.Equal(t, Str, vt1, "type should been Str")
		assert.Equal(t, "hello", v1, "type value should been \"hello\"")

		vt2, v2 := c.Next()
		assert.Equal(t, ListStr, vt2, "type should been ListStr")
		assert.Equal(t, []string{"hello", "world"}, v2, "type value should been \"[hello,world]\"")

		eof, _ := c.Next()
		assert.Equal(t, Eof, eof, "last Next() should return Eof")
	})

	t.Run("Test cParser Mix", func(t *testing.T) {
		c := newCParser("hey,hallo,[hi Freeman,croft]")
		vt1, v1 := c.Next()
		assert.Equal(t, Str, vt1, "type should been Str")
		assert.Equal(t, "hey", v1, "type value should been \"hey\"")

		vt2, v2 := c.Next()
		assert.Equal(t, Str, vt2, "type should been Str")
		assert.Equal(t, "hallo", v2, "type value should been \"hallo\"")

		vt3, v3 := c.Next()
		assert.Equal(t, ListStr, vt3, "type should been ListStr")
		assert.Equal(t, []string{"hi Freeman", "croft"}, v3, "type value should been \"[hi Freeman,croft]\"")

		eof, _ := c.Next()
		assert.Equal(t, Eof, eof, "last Next() should return Eof")
	})

	t.Run("Test cParser example", func(t *testing.T) {
		c := newCParser("Jackie Dawson,14/08/1970,311.667.973-47,[marceneiro, assistente administrativo]\nCory Mendoza,11/02/1957,565.512.353-92,[carpinteiro, marceneiro]")
		vt1, v1 := c.Next()
		assert.Equal(t, Str, vt1, fmt.Sprintf("type should been Str, got %s\n", vt1))
		assert.Equal(t, "Jackie Dawson", v1, "type value should been \"Jackie Dawson\"")

		vt2, v2 := c.Next()
		assert.Equal(t, Str, vt2, fmt.Sprintf("type should been Str, got %s\n", vt2))
		assert.Equal(t, "14/08/1970", v2, "type value should been \"14/08/1970\"")

		vt3, v3 := c.Next()
		assert.Equal(t, Str, vt3, fmt.Sprintf("type should been Str, got %s\n", vt3))
		assert.Equal(t, "311.667.973-47", v3, "type value should been \"311.667.973-47\"")

		vt4, v4 := c.Next()
		assert.Equal(t, ListStr, vt4, fmt.Sprintf("type should been ListStr, got %s\n", vt4))
		assert.Equal(t, []string{"marceneiro", "assistente administrativo"}, v4, "type value should been \"[marceneiro, assistente administrativo]\"")

		vt5, v5 := c.Next()
		assert.Equal(t, Str, vt5, fmt.Sprintf("type should been Str, got %s\n", vt5))
		assert.Equal(t, "Cory Mendoza", v5, "type value should been \"Cory Mendoza\"")

		vt6, v6 := c.Next()
		assert.Equal(t, Str, vt6, fmt.Sprintf("type should been Str, got %s\n", vt5))
		assert.Equal(t, "11/02/1957", v6, "type value should been \"11/02/1957\"")

		vt7, v7 := c.Next()
		assert.Equal(t, Str, vt7, fmt.Sprintf("type should been Str, got %s\n", vt5))
		assert.Equal(t, "565.512.353-92", v7, "type value should been \"\"565.512.353-92\"")

		vt8, v8 := c.Next()
		assert.Equal(t, ListStr, vt8, fmt.Sprintf("type should been ListStr, got %s\n", vt8))
		assert.Equal(t, []string{"carpinteiro", "marceneiro"}, v8, "type value should been \"[carpinteiro, marceneiro]\"")

		eof, v := c.Next()
		assert.Equal(t, Eof, eof, fmt.Sprintf("last Next() should return Eof, got %s: %+v", eof, v))
	})

	t.Run("Test cParser multiline", func(t *testing.T) {
		src := `lindsey craft,19/05/1976,182.845.084-34,[carpinteiro]
Jackie dawson,14/08/1970,311.667.973-47,[marceneiro, assistente administrativo]
Cory mendoza,11/02/1957,565.512.353-92,[carpinteiro, marceneiro]`

		c := newCParser(src)

		for t, _ := c.Next(); t != Eof; {
			t, _ = c.Next()
		}
	})

}

// func TestLoadData(t *testing.T) {

// 	fCandidates, err := os.Open("candidatos.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer fCandidates.Close()

// 	fConcourses, err := os.Open("concursos.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer fConcourses.Close()

// 	candidatesData, err := io.ReadAll(fCandidates)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	concoursesData, err := io.ReadAll(fConcourses)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	candidatesParser := newCParser(string(candidatesData))
// 	concoursesParser := newCParser(string(concoursesData))

// 	for tt, s := candidatesParser.Next(); tt != Eof; {
// 		t.Log(s)
// 	}

// 	for tt, s := concoursesParser.Next(); tt != Eof; {
// 		t.Log(s)
// 	}

// }
