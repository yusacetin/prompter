## Prompter
Create shell prompts in a user friendly LaTeX inspired syntax.

### Building
Simply run `go build prompter.go` to generate the executable named `prompter`.

### Usage
The recommended way to use the program is with the following syntax:

`echo "[input string] | ./prompter"`

The input string should contain commands and attributes. Commands set the appearance modifiers and attributes are the special strings that PS1 parses (I'm using the term "attribute" very loosely here). The program supports all the attributes listed [here](https://www.cyberciti.biz/tips/howto-linux-unix-bash-shell-setup-prompt.html) except for \D{}.

Example:

`echo "\bold{ \red{[} \blue{\u} \red{@} \purple{\h} \S \cyan{\W} \red{]}\green{\$} \space }" | ./prompter`

You can also pass the input string as a command line argument, however you MUST use double slashes. Example:

`./prompter \\bold{ \\red{[} \\blue{\\u} \\red{@} \\purple{\\h} \\S \\cyan{\\W} \\red{]}\\green{\$} \\space }`

The output is the string that needs to be set as the PS1 variable. In these examples it is

`\[\e[1m\]\[\e[31m\][\[\e[0m\]\[\e[1m\]\[\e[34m\]\u\[\e[0m\]\[\e[1m\]\[\e[31m\]@\[\e[0m\]\[\e[1m\]\[\e[35m\]\h\[\e[0m\]\[\e[1m\] \[\e[36m\]\W\[\e[0m\]\[\e[1m\]\[\e[31m\]]\[\e[0m\]\[\e[1m\]\[\e[32m\]$\[\e[0m\]\[\e[1m\] \[\e[0m\]`

You can run `export PS1="[output string]"` to use it in your current session or append it to your shell rc file to make it your default prompt. You can use command substitution like `export PS1="$(echo "\bold{ \red{[} \blue{\u} \red{@} \purple{\h} \S \cyan{\W} \red{]}\green{\$} \space }" | ./prompter)"` to reduce the number of steps. I'm aware of how inconvenient this is and will work on a way to automate it ~when~ if I feel like it.

### List of cCmmands and Custom Attributes
| Command       | Function                                   |
| ------------- | ------------------------------------------ |
| \blackbg{}    | Begin black background color               |
| \redbg{}      | Begin red background color                 |
| \greenbg{}    | Begin green background color               |
| \brownbg{}    | Begin brown background color               |
| \bluebg{}     | Begin blue background color                |
| \purplebg{}   | Begin purple background color              |
| \cyanbg{}     | Begin cyan background color                |
| \whitebg{}    | Begin white background color               |
| \black{}      | Begin black text color                     |
| \red{}        | Begin red text color                       |
| \green{}      | Begin green text color                     |
| \brown{}      | Begin brown text color                     |
| \blue{}       | Begin blue text color                      |
| \purple{}     | Begin purple text color                    |
| \cyan{}       | Begin cyan text color                      |
| \white{}      | Begin white text color                     |
| \bold{}       | Begin bold text                            |
| \dim{}        | Begin dimmed text color                    |
| \italic{}     | Begin italic text                          |
| \underline{}  | Begin underlined text                      |
| \blink{}      | Begin blinking text                        |
| \reverse{}    | Begin inverting text and background colors |
| \overline{}   | Begin overlined text                       |
| \S, \space    | Insert space                               |

### More Example Strings
#### Example commands
* `\bold{ \red{[} \blue{\u} \red{@} \purple{\h} \S \cyan{\W} \red{]}\green{\$} \space }`  
![Screenshot 1](https://yusacetin.github.io/project-screenshots/prompter/1.png)  
* `\bold{\purple{[}} \italic{\t} \bold{\purple{]} \blue{\$}}`  
![Screenshot 2](https://yusacetin.github.io/project-screenshots/prompter/2.png)  
* `\greenbg{\white{\bold{\u \space \$ \space}}}`  
![Screenshot 3](https://yusacetin.github.io/project-screenshots/prompter/3.png)

### The Old Version
See *.old files for the old version.