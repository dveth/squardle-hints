# Squardle Hints

This project is designed to provide hints for the Squardle game, which is available to play at the [Squardle website](https://squaredle.app/).

The project uses the NWL2020 list of words and filters them based on user inputs. The available filters are

- Letters at beginning of word (for example 'te' matching 'test' and 'teeth' but not 'beast')
- Letters word contains (for example 'te' matching 'test' and 'esteem' but not 'beast')
- Letters at end of word (for example 'te' matching 'equate' and 'mate' but not 'beast')
- Length of word
- Letters available in word

The executable will start up a web server which will then make a website available which the user can visit to get words that match the specified filters.
