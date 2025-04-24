// ( ) { } [ ]
//open brackets must be closed by the same type of bracket
// open brackets must be closed in the correct order
// Every close bracket has a corresponding open bracket of the same type


// Use a stack PUSH to a stack if it is an opening one, pop if there is a close one, return "stack is empty"

func isValid(s string) bool {

    stack := []rune{}
    brackets := map[rune]rune{
         ')': '(',
        '}' : '{',
        ']' : '[',
    }
    
    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
        }

        if c == ')' || c == ']' || c == '}' {
            var top rune
            if len(stack) > 0{
                top, stack = stack[len(stack)-1], stack[:len(stack)-1]
            }
            
            if top != brackets[c]{
                return false
            }
        }
    }
    return len(stack) == 0
}
