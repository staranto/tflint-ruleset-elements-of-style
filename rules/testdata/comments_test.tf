# #########
# Tests that will emit issues.

#Bad comment (jammed)
//Bad comment (jammed)

# This comment is way too long and it will definitely extend beyond the eighty character limit that we have set for this rule.

/*
  Block comments
  are not allowed.
*/

resource "foo" "bar" {
  # Indented comment that is also way too long and should trigger the rule because it goes past column 80.
}

# #########
# Tests that will not emit issues.

# Good comment
// Good comment

# Short comment.

