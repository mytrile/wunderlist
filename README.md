# wunderlist

wunderlist is Go wrapper for [Wunderlist](https://www.wunderlist.com/) [unofficial API](https://wunderpy.readthedocs.org/en/latest/wunderlist_api/index.html).

## Installation

      $ go install github.com/mytrile/wunderlist

## Example

Get all available tasks

``` go
package main

import (
  "fmt"
  "github.com/mytrile/wunderlist"
)

func main() {
  client := wunderlist.NewClient("your@email.com", "password")
  tasks, _ := client.Tasks()
  fmt.Printf("%+v", tasks)
}
```

## API

### Account

``` go
// Get the information about the user
account, _ := client.Account()
```

``` go
// Get the settings of the user
settings, _ := client.Settings()
```

``` go
// Get the friends of the user
friends, _ := client.Friends()
```

``` go
// Get the shares of the user
shares, _ := client.Shares()
```

### Tasks

``` go
// Get all available tasks
tasks, _ := client.Tasks()
```

``` go
// Add task
task, _ := client.AddTask("list_id", "title", "true", "2013-10-23T17:06:28Z")
```

### Lists

``` go
// Get all available lists
lists, _ := client.Lists()
```

``` go
// Add task
list, _ := client.AddList("title")
```

### Reminders

``` go
// Get all available reminders
reminders, _ := client.Reminders()
```

``` go
// Add task
reminder, _ := client.AddReminder("task_id", "2013-10-23T17:06:28Z")
```

### Note
Some functions like updating/deleting entities are missing. [Here is the full API](http://godoc.org/github.com/mytrile/wunderlist).

Credits to [bsmt](https://github.com/bsmt) for
[this wonderfull article](http://bsmt.me/blog/2013/03/02/reverse-engineering-the-wunderlist-api/) about reverse engineering the api. You can find the reference API [here](https://wunderpy.readthedocs.org/en/latest/).

### Meta

* Author  - Dimitar Kostov
* Email   - mitko.kostov@gmail.com
* Blog    - <http://mytrile.github.io>
* Twitter - [mytrile](https://twitter.com/mytrile)

## License

(The MIT License)

Copyright (c) 2013 Dimitar Kostov <mitko.kostov@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
