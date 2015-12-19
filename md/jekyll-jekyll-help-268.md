# [Weird error](https://github.com/jekyll/jekyll-help/issues/268)

> state: **open** opened by: **** on: ****

What up Hubbsters, So today i tried to get jekyll working on my computer (64 bit- Widows 7 PC) I downloaded ruby, installed it, and got the dev kit (both 64 bit) I ran them both got them both working and made sure to have exacutables. I made sure to give the devkit its own area to play (i.e C:\devkit )
After this i jumped onto my Command prompt and made sure to have C:\devkit in the prompt.  
Frist I did -
C:\devkit&gt;ruby dk.rb init
that worked completely successful.
Than,
C:\devkit&gt;ruby dk.rb install
that also ran perfect. Here is where the error comes in.
I run,
C:\devkit&gt;gem install jekyll.
It loads for little and than i get this error (I am writing the error verbatim)

ERROR: could not find a valid gem &#x27;jekyll&#x27; &lt;&gt;=0, here is why:
Unable to download data from https://rubygems.org/ - ssl_connect return need=1 errno=0 state=SSLv3 read server certificate B; certificate verify failed ( https://api.rubygems.org/latest_specs.4.8.gz)

Please help if you can. I have a screen shot of it as well if that would be more useful.

Thanks,
Matthew V


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/268#issuecomment-73401384) on: **Sunday, February 08, 2015**

While googling around for this I found a GIST (https://gist.github.com/luislavena/f064211759ee0f806c88) with a couple of strategies and this issue (https://github.com/juthilo/run-jekyll-on-windows/issues/34).

However, I would at first try updating the &#x60;gem&#x60; with &#x60;gem update --system&#x60;.
