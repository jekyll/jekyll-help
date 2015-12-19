# [Trying to install RVM on OSX 10.9.2](https://github.com/jekyll/jekyll-help/issues/33)

> state: **closed** opened by: **** on: ****

I know this isn&#x27;t specifically a Jekyll issue, but I cannot get to the point where I can even install Jekyll.

**rvm install ruby-2.0.0-p451**
&#x60;&#x60;&#x60;
Searching for binary rubies, this might take some time.
Found remote file https://rvm.io/binaries/osx/10.9/x86_64/ruby-2.0.0-p451.tar.bz2
Checking requirements for osx.
Installing requirements for osx.
Updating system....
Installing required packages: libksba.....
Error running &#x27;requirements_osx_brew_libs_install libksba&#x27;,
showing last 15 lines of /Users/Pat/.rvm/log/1399438313_ruby-2.0.0-p451/package_install_libksba.log
++ case &quot;$1&quot; in
++ [[ -t 1 ]]
++ return 1
++ printf %b &#x27;There were package installation errors, make sure to read the log.

Try &#x60;brew tap --repair&#x60; and make sure &#x60;brew doctor&#x60; looks reasonable.

Check Homebrew requirements https://github.com/mxcl/homebrew/wiki/Installation\n&#x27;
There were package installation errors, make sure to read the log.

Try &#x60;brew tap --repair&#x60; and make sure &#x60;brew doctor&#x60; looks reasonable.

Check Homebrew requirements https://github.com/mxcl/homebrew/wiki/Installation
++ case &quot;$_system_version&quot; in
++ return 1
Requirements installation failed with status: 1.
&#x60;&#x60;&#x60;

**brew install libksba**
&#x60;&#x60;&#x60;
Warning: A newer Command Line Tools release is available
Update them from Software Update in the App Store.
==&gt; Downloading ftp://ftp.gnupg.org/gcrypt/libksba/libksba-1.3.0.tar.bz2
Already downloaded: /Library/Caches/Homebrew/libksba-1.3.0.tar.bz2
==&gt; ./configure --prefix=/usr/local/Cellar/libksba/1.3.0
checking for clang option to accept ISO C89... (cached) unsupported
checking dependency style of clang... (cached) none
checking how to run the C preprocessor... clang -E
checking whether clang and cc understand -c and -o together... yes
configure: error: No C-89 compiler found
Configured with: --prefix=/Applications/Xcode.app/Contents/Developer/usr --with-gxx-include-dir=/usr/include/c++/4.2.1

READ THIS: https://github.com/Homebrew/homebrew/wiki/troubleshooting
&#x60;&#x60;&#x60;

**brew doctor**
&#x60;&#x60;&#x60;
Warning: A newer Command Line Tools release is available
Update them from Software Update in the App Store.
&#x60;&#x60;&#x60;

**xcode-select --install**
&#x60;&#x60;&#x60;
Can&#x27;t install the software because it is not currently available from the Software Update server.
&#x60;&#x60;&#x60;

**xcode-select -p**
&#x60;&#x60;&#x60;
/Applications/Xcode.app/Contents/Developer
&#x60;&#x60;&#x60;

**gcc -v**
&#x60;&#x60;&#x60;
Configured with: --prefix=/Applications/Xcode.app/Contents/Developer/usr --with-gxx-include-dir=/usr/include/c++/4.2.1
Apple LLVM version 5.1 (clang-503.0.40) (based on LLVM 3.4svn)
Target: x86_64-apple-darwin13.1.0
Thread model: posix
&#x60;&#x60;&#x60;

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42389642) on: **Tuesday, May 06, 2014**

Hi @pathawks, Did you tried checking the App Store to see if there was an update? I&#x27;m up to date and I have the same version of &#x60;gcc -v&#x60;. Also did you &#x60;brew update&#x60; before installing &#x60;libksba&#x60;? I&#x27;m not a rvm guy so I don&#x27;t know much about it.
---
> from: [**pathawks**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42390206) on: **Tuesday, May 06, 2014**

I just installed Xcode today.
Yes, I tried &#x60;brew update&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42392028) on: **Tuesday, May 06, 2014**

Do you have gcc4.2, from the pre-LLVM days?
---
> from: [**pathawks**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42479394) on: **Wednesday, May 07, 2014**

&gt; Do you have gcc4.2, from the pre-LLVM days?

Perhaps. How can I tell / what can I do about that?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42484113) on: **Wednesday, May 07, 2014**

Can you run &#x60;gcc-4.2 -v&#x60; successfully?

If not, run &#x60;brew install apple-gcc42&#x60; :smiley: 
---
> from: [**pathawks**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42485541) on: **Wednesday, May 07, 2014**

&gt; Can you run &#x60;gcc-4.2 -v&#x60; successfully?

&#x60;&#x60;&#x60;
Using built-in specs.
Target: i686-apple-darwin11
Configured with: /Volumes/Media/Builds/gcc-5666.3/build/obj/src/configure --disable-checking --prefix=/usr --mandir=/share/man --enable-languages=c,objc,c++,obj-c++,fortran --program-transform-name=/^[cg][^.-]*$/s/$/-4.2/ --with-slibdir=/usr/lib --build=i686-apple-darwin11 --program-prefix=i686-apple-darwin11- --host=x86_64-apple-darwin11 --target=i686-apple-darwin11 --with-gxx-include-dir=/include/c++/4.2.1
Thread model: posix
gcc version 4.2.1 (Apple Inc. build 5666) (dot 3)
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42489389) on: **Wednesday, May 07, 2014**

&#x60;&#x60;&#x60;text
configure: error: No C-89 compiler found
&#x60;&#x60;&#x60;

That&#x27;s the problem installing &#x60;libksba&#x60;. Did you find anything via google for that?
---
> from: [**pathawks**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42489780) on: **Wednesday, May 07, 2014**

Google is remarkably unhelpful in this case.

A search for [&quot;*configure:&amp;nbsp;error:&amp;nbsp;No&amp;nbsp;C-89&amp;nbsp;compiler&amp;nbsp;found*&quot;](https://www.google.com/#q=%22configure%3A+error%3A+No+C-89+compiler+found%22) returns **zero** results, while [&quot;*No&amp;nbsp;C-89&amp;nbsp;compiler&amp;nbsp;found*&quot;](https://www.google.com/#q=%22No+C-89+compiler+found%22) just gets me pages of &#x60;configure.ac&#x60; files.
---
> from: [**pathawks**](https://github.com/jekyll/jekyll-help/issues/33#issuecomment-42634571) on: **Thursday, May 08, 2014**

Wiped laptop clean. Did a fresh install of OS X 10.9.
Everything works like a champ now.

I hadn&#x27;t done a fresh install on this laptop for 6 years, so who knows what garbage was getting in the way.

There I fixed it.
