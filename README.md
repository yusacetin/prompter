## Prompter
Create shell prompts in a user friendly LaTeX inspired syntax.

### Build
Simply run `go build prompter.go`.

### Usage
Running `./prompter "[string]"` where `[string]` is your input returns the text that needs to be set as the `PS1` environment variable. You can use `echo export PS1="$(./prompter "[string]")" >> ~/.bashrc` to add it to your bashrc automatically. It reads from stdin when no command line arguments are provided, so you can also do something like `echo "[string]" | ./prompter`.

For the input; use prompt commands (`\u`, `\h`, `\$`, etc) normally, use `\S` or `\space` for space (all blank spaces in the provided input will be ignored), and use commands given in the table below for text formatting.

### List of commands
| Command     | Function                                   |
| ----------- | ------------------------------------------ |
| \blackbg    | Begin black background color               |
| \redbg      | Begin red background color                 |
| \greenbg    | Begin green background color               |
| \brownbg    | Begin brown background color               |
| \bluebg     | Begin blue background color                |
| \purplebg   | Begin purple background color              |
| \cyanbg     | Begin cyan background color                |
| \whitebg    | Begin white background color               |
| \black      | Begin black text color                     |
| \red        | Begin red text color                       |
| \green      | Begin green text color                     |
| \brown      | Begin brown text color                     |
| \blue       | Begin blue text color                      |
| \purple     | Begin purple text color                    |
| \cyan       | Begin cyan text color                      |
| \white      | Begin white text color                     |
| \bold       | Begin bold text                            |
| \dim        | Begin dimmed text color                    |
| \italic     | Begin italic text                          |
| \underline  | Begin underlined text                      |
| \blink      | Begin blinking text                        |
| \reverse    | Begin inverting text and background colors |
| \overline   | Begin overlined text                       |
| \clear      | Clear all formatting                       |
| \S, \space  | Insert space                               |


### Examples
#### Example commands
* `\bold \red [ \blue \u \green @ \purple \h \S \cyan \W \red ] \white \$ \clear \S`
* `\bold \purple [ \clear \italic \t \clear \bold \purple ] \blue \$ \clear \S`
* `\greenbg \bold \white \u \space \$ \space \clear`
#### Export to current session
* `export PS1="$(./prompter '\bold \red [ \blue \u \green @ \purple \h \S \cyan \W \red ] \white \$ \clear \S')"`
* `export PS1="$(./prompter '\bold \purple [ \clear \italic \t \clear \bold \purple ] \blue \$ \clear \S')"`
* `export PS1="$(echo '\cyan \$ \clear' | ./prompter)"`
#### Misc
* Given that `string.txt` contains `\purple \h \S \bold \white \$ \clear \S`,  
  `cat string.txt | ./prompter`  
  To automatically add to bashrc,  
  `echo "PS1='$(cat string.txt | ./prompter)'" >> ~/.bashrc`

### License
GNU General Public License version 3 or later.