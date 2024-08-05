# ascii-art-color

## Objectives
Same as in the ascii-art project but this time with colors.

The output should manipulate colors using the flag --color=<color> <letters to be colored>, in which --color is the flag and <color> is the color desired by the user and <letters to be colored> is the letter or letters that you can chose to be colored. 

    You can choose between coloring a single letter or a set of letters.
    If the letter is not specified, the whole string will be colored.
    The flag must have exactly the same format as above, any other formats will return the following usage message:
        Usage: go run . [OPTION] [STRING]
        EX: go run . --color=<color> <letters to be colored>


## Allowed packages
    Only the standard Go packages.

    Key features include:
    Converting plain text into ASCII art using a predefined standard.txt file.
    Allowing users to highlight specific words or letters in various colors using a --color=<color> flag, supporting colors such as red, green, yellow, blue, purple, cyan, and white.
    Exact Word Matching: Ensures only the exact words/letters specified by the user are colorized, maintaining the integrity of the ASCII art for other text segments.
    Designed to be intuitive and easy to use, requiring minimal input while offering versatile output customization.


## Usage 
###
    With below command the output will be in red 
    $ go run . --color=red "hello"
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $


    With below command both words will be in displayed in blue 
    $ go run . --color=blue "hello world"
     _              _   _                                           _       _
    | |            | | | |                                         | |     | |
    | |__     ___  | | | |   ___         __      __   ___    _ __  | |   __| |
    |  _ \   / _ \ | | | |  / _ \        \ \ /\ / /  / _ \  | '__| | |  / _` |
    | | | | |  __/ | | | | | (_) |        \ V  V /  | (_) | | |    | | | (_| |
    |_| |_|  \___| |_| |_|  \___/          \_/\_/    \___/  |_|    |_|  \__,_|



    While this will colorize in yellow both 'o' in 'hello world' 
    $ go run . --color=yellow "o" "hello world"
     _              _   _                                           _       _
    | |            | | | |                                         | |     | |
    | |__     ___  | | | |   ___         __      __   ___    _ __  | |   __| |
    |  _ \   / _ \ | | | |  / _ \        \ \ /\ / /  / _ \  | '__| | |  / _` |
    | | | | |  __/ | | | | | (_) |        \ V  V /  | (_) | | |    | | | (_| |
    |_| |_|  \___| |_| |_|  \___/          \_/\_/    \___/  |_|    |_|  \__,_|



    This will colorize 'hello world' in green     
    $ go run . --color=yellow "hello world" "hello world"
     _              _   _                                           _       _
    | |            | | | |                                         | |     | |
    | |__     ___  | | | |   ___         __      __   ___    _ __  | |   __| |
    |  _ \   / _ \ | | | |  / _ \        \ \ /\ / /  / _ \  | '__| | |  / _` |
    | | | | |  __/ | | | | | (_) |        \ V  V /  | (_) | | |    | | | (_| |
    |_| |_|  \___| |_| |_|  \___/          \_/\_/    \___/  |_|    |_|  \__,_|