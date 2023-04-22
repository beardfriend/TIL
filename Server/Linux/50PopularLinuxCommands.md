https://www.youtube.com/watch?v=ZtqBQ68cfJc를 보고 정리하였습니다.

강의가 정리된 곳 : https://www.freecodecamp.org/news/the-linux-commands-handbook/#the-linux-man-command

# why use terminal

### whoami

print the user name currently logged in to the terminal session

### man

man is short for manual

```bash
man <command>
```

### clear

```bash
clear -x
ctrl + l
```

### pwd

whenever you feel lost in the filesystem call the `pwd`
command to know where you are

pwd is short for print working directory

### ls

ls is short for list

inside a folder you can list all thie files that the folder contains using the `ls` command

if you add a folder name or path, it will print that folder contents:

```bash
ls /bin
```

### cd

cd is short for change directory

### mkdir

is short for make directory

nested

```bash
mkdir -p /test/ss
```

### touch

### rmdir

### rm

remove

-r recursive
-v view
-f force

### open

ubuntu: `xdg-open <path>`

### mv

is short for move

Once you have a file, you can move it around using `mv` command. You specify the file current path,
and its new path:

```bash
touch pear
mv pear new_pear
```

The `pear` file is now moved to `new_pear`. This is how you rename files and folders

If the last parameter is a folder, thie file located at the first parameter path is going to be moved into that folder.

In this case, you can specify a list of files and they will all be moved in the folder path indentified by the last parameter

### cp

is short for copy

-r

recursive

### head

print file at head

### tail

print file at tail

-f follow

### redirecting standard output

```bash
date >> Some.txt
```

`>` new write

`>>` update

### cat

entire file

```bash
cat file
```

print the content of multiple files

```bash
cat file file2
```

```bash
cat file file2 file3 > everything.txt
```

`cat -n` line number

### less

It shows you the content stored inside a file, in a nice and interactive UI

space -> one page
/text search
g start line of file
Shift + g end line of file

### echo

The echo command does one simple job: it prints to the output the argument passed to it.

`echo "test" > file.txt`

add or create small things file quickly

echo after text redirect into the file

### wc

word count

`number of line` `number of word` `number of bytes`

wc –> word, line, character, and byte count

### piping

`|` piping

### sort

suppose you have a text file which contains the names of dogs

its not change the file, simply sort output

i have trying sort nums, it doesnt numeric sort

`-n` option sort numericly

`-r` option reverse

`-u` unique number sorted

example
`cat butcher.txt groceries.txt | sort`

### uniq

unique is a command useful to sort lines of text

adjacent(인접한)단어만 중복제거

`-u`

print only unique

`-d`

print only duplicated

`-c`

print count duplicated

### expansion

~ = tildy = home

`echo '~'` -> `~`
$PATH 
$USER = current user name

-   = everything

? = one char

`echo *.??` -> fasf.md qwe.go ...

`echo *.???` -> asd.tsx

{} combination

`echo {a,b,c}.txt` -> a.txt b.txt c.txt

`echo {1..5}` -> 1, 2, 3, 4, 5

### diff

diff is a handy command. Suppose you have 2 files, which contain almost the smae information,
but you can'nt find the difference between the two.

diff will process the files and will tell you what's the difference.

`-u` like git

24d25 first file 24line **delete** secondfile 25line
24c25 **change**
24a25 **added**
