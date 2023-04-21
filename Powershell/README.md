- [Printing Messages](#printing-messages)     
- [Powershell Variables](#powershell-variables)    
- [String Operations](#string-operations)    
- [if else statements](#if-else-statements)  
- [Powershell Operators List](powershell-operators-list)
- [For and while loops](#for-and-while-loops)    
- [Functions](#functions)  
- [Array and Operations](#array-and-operations)  
- [Hash Tables](#hash-tables)  
- [Exception Handling](#exception-handling)   
- [Powershell Comments](#powershell-comments)  
- [Powershell Providers and Modules](#powershell-providers-and-modules) https://www.computerworld.com/article/3154158/all-about-powershell-providers-and-modules.html  
- [Using Objects in Powershell](#using-objects-in-powershell) https://www.computerworld.com/article/2954261/understanding-and-using-objects-in-powershell.html


### Printing Messages 

__Print Hello World__  

`Write-Output` command is used to print string data on console. Example : 

```  
Write-Output "Hello World"
```  

Printing Numbers 

```  
Write-Output 12.34 
Write-Output 15420
```  

__Alias for `Write-Output`__    

The `Write-Output` has two aliases aleardy set `echo` and `write` which is easier to use.  

```  
echo "Hello World"
write "Hello World Again"  
```  

__Printing Variable values__  

```  
$num = 12345
write $num

$num1 = 12.673
write $num1

$name = "Don Joe"
write "Hello $name"

$var = 'X'
write $var
```  

__Taking User Input__

`Read-Host` command is used to take input from user.  

```
$name = Read-Host "Enter your name "
write $name
```  

Take Password as input 

```  
$pass = Read-Host 'Enter password ' -AsSecureString
```  

Printing password as normal text user 

```  
[Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($pass))
```  

__Printing colored Values :__  

`Write-Host` command prints colourful output.  

```  
write-host "Hello world" -ForegroundColor green
```  

Changing background color  

```  
Write-Host "Hello world" -ForegroundColor darkgreen -BackgroundColor white
```  

## Powershell Variables 

PowerShell variables are named objects. As PowerShell works with objects, these variables are used to work with objects. Powershell variables have dollor sign `$` around them. Example :  

```  
$num = 12450
$name = "Done Joe" 
```   

Storing multiline string into variable 

```  
$str = @"
My name is vignesh
Am from chennai
Am a freelancer
"@
```  

Story multiple values in comma seperated list    

```  
$lists = 10,20,30,40,50,100
```   

Printing values by index 

``` 
echo $lists[0]
echo $lists[1]
echo $4lists[2]
echo $ists[-1]
```  

Counting total number of values 

```  
echo $lists.count
``` 

### String Operations  

__String Concatenation :__  

Suppose we have three string  

```  
$str1 = "Hello world"
$str2 = "This is test"
```  

Concatnating two strings 

Method 1  

```  
> $str3 = $str1 + $str2
> echo $str3 

Hello worldThis is test
```   

Method 2   

```  
> $str3 = $str1 + " " + $str2
> echo $str3

Hello world This is test
```    

Method 3  

```  
> $str3 = "{0} {1}" -f $str1,$str2  
> echo $str3 

Hello world This is Text
```  

Method 4  

```  
> $str3 = $str1,$str2 -join " "
> echo $str3 

Hello world This is Text
```  

Method 5  

```  
> $str3 = [System.String]::Concat($str1," ",$str2)
> echo $str3 

Hello world This is Text
```

Method 6  

```  
> $str3 = "First String ${str1} : Second String ${str2}" 
```

Concatenating multiline Strings  

```  
> $str1 = @"
Hello world
This is First String
"@

> $str2 = @"
Hello world Again 
This is Second String
"@

> $str3 = $str1,$str2 -join "`n"
> echo $str3

Hello world
This is First String
Hello world Again 
This is Second String
```   

__Replacing String :__  

Using `replace()` function. Syntax 

```  
$MyStr.replace(OldString, NewString)  
```  

Example :  

```  
> $MyStr = "This is the sample String"
> echo $MyStr.replace("sample", "Test")

This is the Test String  

> echo $MyStr.replace("sample", "Test").replace("the", "a")

This is a Test String
```   

__Splitting String :__  

Using `split()` function  

```  
> $str = "Hello world This is Test String"
> echo $str.split(" ")  

Hello
world
This
is
Test
String
```  

__Comparing String :__   

* `CompareTo()` function : Returns 0 if matched, 1 if some parts matched, -1 if nothing matched.  

```   
> $str = "Test"
> $str.CompareTo("Test")
0 

> $str = "Test String"
> $str.CompareTo("Test")
1 

> $str = "Test String"
> $str.CompareTo("Hello")
-1
```   

* `Equals()` function : Retrns True if matched, otherwise false.   

```  
> $str = "Test String"
> $str.Equals("Test String")
True  

> $str = "Test String"
> $str.Equals("Test")
False 
```   

__Substring function :__ 

`.Substring()` function prints/returns a particular part of a string defined by start index and lenth form starting index.  

```   
$str.Substring(StartIndex, lengthOfString)  
```  

Example :  

```  
> $str = "This is a Sample test String"
> $str.Substring(0,10)

This is a

> $str.Substring(10, 12)

Sample test
```  

### If Else Statements  

__if else Statement :__  

Syntax :  

```  
if (condition) {
    // Statement 
} else {
    // Statement
}
```  

Example : 

```  
$x = 10

if ($x -le 15) {
    Write-Output "x is less then or equal to 15"
} else {
    Write-Output "x is greater then 15"
}
```  

__if elseif Statement :__  

Syntax :  

```   
if (condition) {
    // Statement 
} elseif (condition) {
    // Statement
} else {
    // Statement 
}
```  

Example : 

```  
$wind = Get-Service WinDefend

if ($wind.Status -eq "Running") {
    write "Windows Defender is running"
} elseif ($wind.Status -eq "Stopped") {
    write "Windows Defender is Stopped"
} else {
    write "Windows Defender is in Start or Stopped pending status"
}
``` 

### Powershell Operators List 

__Arithmetic Operator :__  

|Operator|Description|Example|  
|---|---|---|  
|+ (Addition)|Add two or more operand|A+B|  
|- (Subtraction)Subtract two or more operand|A-B|  
|* (Multiplication)|Multiply two or more operand|A*B|  
|/ (Division|Divides left-hand operand by right hand operand|A/B|  
|% (Modulus)|Divides left-hand operand by right hand operand and returns remainder|A%B|  

__Logical Operators :__

|Operator|Description|Example|  
|---|---|---|
|AND (logical and)|If both the operands are non-zero, then the condition becomes true.| A -AND B|  
|OR (logical or)|If any of the two operands are non-zero, then the condition becomes true.|A -OR B|  
|NOT (logical not)|If a condition is true then Logical NOT operator will make false.|-NOT A, -NOT(A -AND B)|  

__Comparision Operators :__  

|Operator|Description|Example|   
|---|---|---|  
|eq (equals)|Compares two values to be equal or not.|A -eq B|  
|ne (not equals)|Compares two values to be not equal|A -ne B|  
|gt (greater then)|Compares first value to be greater than second one.|A -gt B|  
|ge (greater then or equals to)|Compares first value to be greater than or equals to second one.|A -ge B|  
|lt (less then)|Compares first value to be less than second one.|A -lt B|   
|le (less then or equals to|Compares first value to be less than or equals to second one.|A -le B|   

__Assignment Operators :__  

|Operator|Description|Example|  
|---|---|---|
|= (Assignment operator)|Assigns values from right side operands to left side operand.|C = A + B|   
|+= (Add and Assignment operator)|It adds right operand to the left operand and assign the result to left operand.|C += A| 
|-= (Subtract and Assignment operator)|It subtracts right operand from the left operand and assign the result to left operand.|C -= A|
|*= (Multiply and Assignment operator)|It mulltiply right operand from the left operand and assign the result to left operand.|C *= A|  

__Increment And Decrement Operators :__   

|Operator|Description|Example|  
|---|---|---|  
|++X|Pre-increment|++A|  
|X++|Post-increment|A++|  
|--X|Pre-decrement|--A|  
|X--|Post-decrement|A--|  

### For And While Loops  

Suppose we have a variable 

```  
$nums = 10,20,30,40,50,60,100
```  

__For loop :__  

```  
for($i = 0; $i -lt $nums.length; $i++) {
    write $nums[$i]
}
```

__foreach loop :__  

```  
foreach($item in $nums) {
    write $item
}
``` 

__while loop :__  

```  
$c = 0
while ($c -lt $nums.length) {
    write $nums[$c]
    $c++
}
```  

__do while loop :__  

```  
$c = 0
do {
    write $nums[$c]
    ++$c
} while($c )
```  

### Functions  

__Function structure :__    

```    
function FunctionName(param1, param2...) {
    // function Statement
    return Value
}
```    

example :  

```    
function TestFunc($str) {
    write $str
}
```   

Calling a function    

```    
TestFunc "hello world"
```    

__Another way to defininig parameters :__  

```   
function TestFunc {
    param (
        $ComputerName
    )
    write $ComputerName
}
```  

__Function with return Value :__  

```   
function TestFunc($num1, $num2) {
    return $num1 + $num2;
}

TestFunc 10 200
```   

For more Details on Powershell Functions visit Link : https://docs.microsoft.com/en-us/powershell/scripting/learn/ps101/09-functions?view=powershell-7.1  
### Array and Operations  


https://www.educba.com/array-in-powershell/?source=leftnav  

__Defining an Array on Powershell :__  

Method 1 : By using `@` operator  

```  
$MyArr = @('Hello', 'World', 'This', 'is', 'Test')
```   

Accessing elements 

with index 

```  
$MyArr[0]
$MyArr[1]
$MyArr[2]
$MyArr[3]

// Accessing last element 
$MyArr[-1]
```  

with loops 

```  
foreach ($elem in $MyArr) { echo $elem }
```  

Method 2 :  

```
$MyArr = 'Hello','World','This','is','Test'
```  

Creating Numerical array :  

```   
$nums = 1,2,3,4,5,10,20,30
```  

another way  

```  
[int[]] $nums = 1,2,3,4,5
```  

__Finding length of a string or an array :__  

`$Array.length` gives the length of array    

```  
$MyStr = 1,2,3,4,5,10

$MyStr.length
```  

Example :  

``` 
for ( $i=0; $i -lt $MyStr.length; $i++ ) {
    echo $MyStr[$i];
}
```  

Adding, updating, removing items from ArrayList :__   

Array does not allow to add new element, but by using arraylist 

```  
$MyStr = [System.Collections.ArrayList]::new()
$MyStr.Add(10)
$MyStr.Add(20)
$MyStr.Add(30)
$MyStr.Add(40)
```  

Remove An element from arraylist `$MyStr.Add('element') 

```  
$MyStr.Remove(30)
```

### Hash Tables 

Hash Table is basically structured array with key, value pair. It is also known as Dictonary or Associative array.   

```   
$MyDict = @{key1="Value1"; key2="Value2"; key3="Value3"}
```  

Accessing Elements 

``` 
$MyDict["key1"] // gives the values 
```  

Listing all the keys  

```  
$MyDict.keys
```

Listing all the values 

```  
$MyDict.values  
```

Count number of key,value pairs  

```  
$MyDict.count 
```  

Add new item on ArrayList  

```  
$MyDict.Add("key", "value")
```  

Remove item from Arraylist  

```   
$MyDict.Remove("Kkey_Name")
```  

### Exception Handling  

Structure of Exception Handling on Powershell  

```   
try {
    // Statements
    // ...
} catch {
    // Statements
    // ...
}
```

Example :

```   
try {
    SomeERRORCommnd
} catch {
    echo "Command not found"
}
```

You can also use Final block

```
try {
    // Statements
    // ...
} catch {
    // Statements
    // ...
} final {
    // Statements
    // ...
}
```

For more details visit here : https://docs.microsoft.com/en-us/powershell/scripting/learn/deep-dives/everything-about-exceptions?view=powershell-7.1


## Powershell Comments  

Single line comments  

```  
# This is a commant 
echo "Hello world"
```  

Multiline comment  

```    
<# 
Hello world 
This is the test
comment line
#>  
echo "Hello world"
```    
