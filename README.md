## Coding Challenge #50 - Build Your Own Xargs

This challenge is to build your own version of the Unix command line tool `xargs`!

The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Writing simple parts connected by clean interfaces - each tool does just one thing and provides a simple CLI that handles text input from either files or file streams.
    
- Design programs to be connected to other programs - each tool can be easily connected to other tools, via files and streams, to create incredibly powerful compositions.
    

Following these philosophies has made the Unix command line tools some of the most widely used software engineering tools which can be chained together to create far more complex and powerful set of tools that you’d expect.

Xargs epitomises the philosophy by providing a tool to allow us to connect together programs effectively, using the output of one to configure the behaviour of the next.

### If You Enjoy Coding Challenges Here Are Four Ways You Can Help Support It

1. Refer a friend or colleague to the newsletter. 🙏
    
2. Sign up for a [paid subscription](https://substack.com/redirect/40b5b13b-ebf1-46ee-9d12-1c0dbbc2a32b?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) - think of it as buying me a coffee ☕️ twice a month, with the bonus that you also get 20% off any of my [courses](https://substack.com/redirect/e3eccb77-b49a-475e-87d5-e3edf75ffdaf?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4).
    
3. Buy one of my [courses](https://substack.com/redirect/e3eccb77-b49a-475e-87d5-e3edf75ffdaf?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) that walk you through a Coding Challenge.
    
4. If you work for a company that sells to software engineers, encourage them to [sponsor the newsletter](https://substack.com/redirect/8c771979-ab4d-4e87-acd1-6028c3fab5ed?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4). 🙏
    

## The Challenge - Building You Own Xargs

This challenge is to build your own version of `xargs`. As always with command line tools a great way to find out what the tool does and how to use it is to use `man`:

```
NAME
     xargs – construct argument list(s) and execute utility

SYNOPSIS
     xargs [-0oprt] [-E eofstr] [-I replstr [-R replacements]
           [-S replsize]] [-J replstr] [-L number] [-n number [-x]]
           [-P maxprocs] [-s size] [utility [argument ...]]

DESCRIPTION
     The xargs utility reads space, tab, newline and end-of-file
     delimited strings from the standard input and executes utility
     with the strings as arguments.

     Any arguments specified on the command line are given to
     utility upon each invocation, followed by some number of the
     arguments read from the standard input of xargs.  This is
     repeated until standard input is exhausted.
```

You can read about how useful `xargs` can be in my [Developing Skills](https://substack.com/redirect/5ad63aa1-dc51-4b3e-ab55-7bf17ffcd89d?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)newsletter article that explains how I used it to [build a simple load testing tool for a RESTful API](https://substack.com/redirect/e07a76bd-cc85-4319-abb6-307d017a06b3?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4).

### Step Zero

In this introductory step you’re going to set your environment up ready to begin developing and testing your solution.

I’ll leave you to setup your IDE / editor of choice and programming language of choice.

### Step 1

In this step your goal is to build the command `ccxargs` that will take a whitespace separated set of strings from standard in and convert them into command line arguments that can be passed to a command (referred to as `utility` in the man page quoted above).

You can test your code using, this command below to create three text files we can use for testing:

```
% for i in {1..3}; do echo "This is file ${i}" > test-${i}.txt; done;
```

Then in the same directory we can use `ls` to create a whitespace separated list of files and pipe that into our `ccxargs` program which we will tell to run the command `cat` with each of the items in the list as the argument to `cat`:

```
% ls | ccxargs cat                                                     
This is file 1
This is file 2
This is file 3
```

This is the equivalent of having done:

```
% cat test-1.txt test-2.txt test-3.txt 
This is file 1
This is file 2
This is file 3
```

Bonus points if you use your own version of `cat` from the [build your own cat](https://substack.com/redirect/2bea0fa7-de2d-4c3d-9e1b-c447dc9a3b3b?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)Coding Challenge.

### Step 2

In this step your goal is to support additional command line arguments being passed to the utility command. We can test this using `sed` like so:

```
% ls | ccxargs sed s/file/test/g                                    
This is test 1
This is test 2
This is test 3
```

But don’t stop there, ensure you test with 2+ arguments to the command too.

### Step 3

In this step your goal is to support the `-n` option, which does this (from the `man` page):

```
 -n number, --max-args=number
             Set the maximum number of arguments taken from standard
             input for each invocation of utility.  An invocation of
             utility will use less than number standard input arguments
             if the number of bytes accumulated (see the -s option)
             exceeds the specified size or there are fewer than number
             arguments remaining for the last invocation of utility.  The
             current default value for number is 5000.
```

You can then test this like so:

```
% ls | ccxargs -n 1 cat                 
This is file 1
This is file 2
This is file 3
```

This time your code should have invoked `cat` three times, one after the other.

### Step 4

In this step your goal is to support the `-P` option, which does this (from the `man` page):

```
     -P maxprocs, --max-procs=maxprocs
             Parallel mode: run at most maxprocs invocations of utility
             at once.  If maxprocs is set to 0, xargs will run as many
             processes as possible.
```

To test this I suggest creating a text file with a list of URLs in it, say `urls.txt`, then use `ccxargs` to invoke `curl` to download the pages.

```
% cat urls.txt | ccxargs -n 1 -P 1 curl
```

Which will dump the content of the websites you hit to your console. You might like to time it.

Then run the test again with a higher value of P to see the overall time reduced as the `curl` requests are sent concurrently.

```
% cat urls.txt | ccxargs -n 1 -P 10 curl
```

### Going Further

You can take this challenge further by adding support for the other options detailed on the man page, or visiting the [build your own cat](https://substack.com/redirect/2bea0fa7-de2d-4c3d-9e1b-c447dc9a3b3b?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) and [build your own curl](https://substack.com/redirect/fd1988dd-30f7-4a38-bc7c-377ed9633d7c?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) challenges and implementing and using them.

## **Featured Jobs:**

This is a new section of Coding Challenges where I’ll feature a couple of open roles with the aim of introducing you as a great candidate to a great new role. Here are three I thought looked interesting:

- [Senior Full-Stack Software Engineer at Gynger (US - Remote)](https://substack.com/redirect/b2aa122a-ce28-4b6f-8cab-7b4e42023e82?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)
    
- [Senior Full Stack Engineer at Rec (US - Remote)](https://substack.com/redirect/752bb802-e6cf-40f7-80f2-77a61e52067d?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)
    
- [Staff Software Engineer, User Data Team at OneSignal (US - Remote)](https://substack.com/redirect/77c1141e-01cf-4435-8ec9-0c488a9565fb?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)
    

### 2 Other Ways I Can Help You:

1. I write another FREE newsletter **[Developing Skills](https://substack.com/redirect/5ad63aa1-dc51-4b3e-ab55-7bf17ffcd89d?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)** that helps you level up the other skills you need to be a great software developer.
    
2. I have some courses available:
    
    1. **[Build Your Own Redis Server (Python Edition)](https://substack.com/redirect/c8b2ec6c-77f3-4aca-b96e-2146755628bb?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)** which guides you through solving the Redis Coding Challenge in Python.
        
    2. **[Build Your Own Shell (Go Edition)](https://substack.com/redirect/0e493be9-77f4-46bb-9950-5e1e69fa19a1?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)** which guides you through solving the Shell Coding Challenge in Go.
        

### Share Your Solutions!

If you think your solution is an example of the developers can learn from please share it, put it on GitHub, GitLab or elsewhere. Then let me know via [Twitter](https://substack.com/redirect/a89aa3fe-aee5-4550-9591-93a5e11223de?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) or [LinkedIn](https://substack.com/redirect/7fc3de1c-8a4d-4fbb-a437-d5990286c0a7?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) or just post about it there and tag me.

### Request for Feedback

I’m writing these challenges to help you develop your skills as a software engineer based on how I’ve approached my own personal learning and development. What works for me, might not be the best way for you - so if you have suggestions for how I can make these challenges more useful to you and others, please get in touch and let me know. All feedback greatly appreciated.

You can reach me on [Twitter](https://substack.com/redirect/a89aa3fe-aee5-4550-9591-93a5e11223de?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4), [LinkedIn](https://substack.com/redirect/7fc3de1c-8a4d-4fbb-a437-d5990286c0a7?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4) or through [SubStack](https://substack.com/redirect/e466cb7a-3663-4e88-8e7a-2d934a47afe9?j=eyJ1IjoiNjBlb2MifQ.8rn2BGPA17vWEgO4Xc0K6qwyLo9ABhcxMkO5h3PoZK4)

Thanks and happy coding!

John