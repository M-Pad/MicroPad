package lua

import lua "github.com/yuin/gopher-lua"

// Defines a plugin's details and actions.
type PluginMeta struct {
	// Plugin details
	Name        string
	Version     string
	Description string
	Homepage    string
	Source      string

	// Plugin "Actions"
	Hooks    map[string]lua.LFunction
	Commands []EditorCommand
}

/*
Defines an editor command and completion.

The completion function is called every keystroke after the user invokes
the completion keybind ( TAB ). The function returns plausible suggestions
based of the current input ( as a string ) or an empty string if the user
hasn't provided any arguments.

Examples:

	{
	  Name = "HelloWorld",
	  Description= "Puts a friendly message in the console",
	}

	{
	  Name = "ComplexCommand",
	  Description = "A really, really complicated command",

	  Completion = function (arguments)
	    print(arguments) -- string, e.g. hello world
	    return response
	  end
	}
*/
type EditorCommand struct {
	Name        string
	Description string

	Completion *lua.LFunction
}
