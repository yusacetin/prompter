## Prompter
Create shell prompts in a user friendly LaTeX inspired syntax.

### Build
Simply run `go build prompter.go`.

### Usage
Running `./prompter [string]` where `[string]` is your input returns the text that needs to be set as the `PS1` environment variable. You can use `echo "export PS1='$([string])'" >> ~/.bashrc` to add it to your bashrc automatically. It reads from stdin when no command line arguments are provided, so you can also do something like `echo "[string]" | ./prompter`.

For the input; use prompt commands (`\u`, `\h`, `\$`, etc) normally, use `\S` or `\space` for space (all blank spaces in the provided input will be ignored), use one of the commands in the maps defined in prompter.go to begin a format, and use `\clear` to clear formatting.

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