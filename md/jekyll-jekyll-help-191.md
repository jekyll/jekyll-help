# [One blank line in the code block when using highlight.js](https://github.com/jekyll/jekyll-help/issues/191)

> state: **closed** opened by: **** on: ****

I changed my syntax highlighter to highlight.js. It works, but there is always one blank line above the code . I can&#x27;t find where to adjust it. highlight.js&#x27;s CSS is:

.hljs {
  display: block;
  overflow-x: auto;
  padding: 0.5em;
  background: #23241f;
  border-radius: 8px;
  -webkit-text-size-adjust: none;
  max-width:  600px;
  font-size: 0.8em;
}

![blanklineissue](https://cloud.githubusercontent.com/assets/1506580/4967714/d311c514-6820-11e4-8920-91a7cf26d9c6.png)


### Comments

---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62347556) on: **Sunday, November 09, 2014**

What does the generated HTML look like?
---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62347663) on: **Sunday, November 09, 2014**

@mhulse it&#x27;s here: http://legato.ninja
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62349011) on: **Sunday, November 09, 2014**

Looks like you got a rogue ~~&#x60;p&#x60; tag~~ line return in there. Check it:

![p-tag-pre](https://cloud.githubusercontent.com/assets/218624/4972314/6689dfc4-68a9-11e4-8bab-d9d4c5d46497.gif)

What does your raw post input look like?
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62349370) on: **Sunday, November 09, 2014**

&gt; rogue &#x60;p&#x60; tag

I meant, line return.
---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62356810) on: **Monday, November 10, 2014**

@mhulse Wow~It&#x27;s very impressive to have animated solution, cool. I am not sure about what is &#x60;rouge p tag&#x60;. Here is my raw post input.

![raw_post](https://cloud.githubusercontent.com/assets/1506580/4973212/b71f02e4-68b7-11e4-838b-224fbf764119.png)

---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62359887) on: **Monday, November 10, 2014**

Thanks. Replying from my phone. 

Try moving the pre and code tags on same line as the code sample. This should take care of your space. It&#x27;s not pretty looking, but would tighten up the visual gap. 

Check to see if highlight.js has a trim L/R feature.  But I think your best bet is to apply my suggested fix above.
---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62361245) on: **Monday, November 10, 2014**

@mhulse It seems I found a solution here, but still don&#x27;t know why. Solution is to write the code snippet follows &#x60;&lt;code&gt;&#x60;tag without &#x60;new line&#x60;

![solution](https://cloud.githubusercontent.com/assets/1506580/4973786/961189b6-6900-11e4-9fe7-c3b7a4633209.png)

It fixes the problem, but makes raw post not readable in terms of code snippet.

---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62362608) on: **Monday, November 10, 2014**

@mhulse Thank you for your kind help. Yes, it(first solution) do fix this problem! I guess I need to search and learn what is &#x60;trim L/R feature&#x60;. :)
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62362817) on: **Monday, November 10, 2014**

&gt; It seems I found a solution here, but still don&#x27;t know why.

Whitespace inside &#x60;&lt;pre&gt;&#x60; elements is always respected by the browser, and output as-is. A newline is treated the same as any other whitespace character.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62365754) on: **Monday, November 10, 2014**

Glad to help. Trim is white space removal at beginning and end of strings. You can trim Right (trim R), Left (trim L), or both (trim) depending on Lang/function used. It&#x27;s a long shot, but maybe highlight JS has an option for trimmin? Since your using JS anyway, you could prob write your own JS trim on the code/pre blocks.  OTOH, like has been mentioned, white space is respected, so I think your best bet is to cram the lines or use fenced code and the jekyll highlighter.  Personally, iirc, I&#x27;ve always crammed lines when I&#x27;ve run into this issue. Unfortunately it&#x27;s been a while since I&#x27;ve thought about this scenario, so take what I&#x27;m saying with a grain of salt. :)
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62366448) on: **Monday, November 10, 2014**

This could work too (not tested):

&#x60;&#x60;&#x60;
&lt;pre&gt;&lt;code
&gt;line of code
Here and ...
Here
&lt;/code&gt;&lt;/pre&gt;
&#x60;&#x60;&#x60;

Not sure if that looks any better though. :D

Sometimes I do similar technique as above to kill white space from inbetween inline block elements. ;)
---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62380129) on: **Monday, November 10, 2014**

@mhulse Thank you for your help. I have a happy day.
---
> from: [**legatoo**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-62380194) on: **Monday, November 10, 2014**

@nternetinspired Thank you! Now I know it.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-63563704) on: **Tuesday, November 18, 2014**

Love the love! Thanks for jumping in, @mhulse!
---
> from: [**bryc**](https://github.com/jekyll/jekyll-help/issues/191#issuecomment-89039983) on: **Thursday, April 02, 2015**

&#x60;&#x60;&#x60;js
var y = document.querySelectorAll(&quot;pre code&quot;);
for(var i = 0; i &lt; y.length; i++) {
  y[i].innerHTML = y[i].innerHTML.replace(&quot;\n&quot;, &quot;&quot;);
}
&#x60;&#x60;&#x60;
Removes the first newline from each code block.
