# [documenting a bit more the parsing of yaml with liquid](https://github.com/jekyll/jekyll-help/issues/266)

> state: **open** opened by: **** on: ****

First, I&#x27;m new to jekyll, and I think it&#x27;s great.

I&#x27;m struggling a bit with the parsing of yaml files using for loops.

My understanding is that there are basically two types of objects that one might find in yaml files: lists and so-called &quot;associative arrays&quot; that are some kind of dictionaries. These two structures can be nested, and basically that is all there is to it. (See http://en.wikipedia.org/wiki/YAML)
 
I find it hard to guess the exact liquid syntax that can be used in jekyll as how to reach a particular element, and I find that the online documentation (which is great for most purposes) is a bit too scarce, unless I have missed some part of it.

So, my suggestion would be to:
1. add more examples of how to parse a yaml file using liquid.
2. give references to some kind of place where the types used in liquid are explained more precisely.
This could be done by expanding the section:
http://jekyllrb.com/docs/datafiles/#the-data-folder
or by adding a new section.

I understand that this is borderline between liquid and jekyll, but it&#x27;s such an important practical step to use the full power of jekyll that it definitely would help new users a lot if more details were given.

To make things concrete, here is a problem for which I only have an ugly solution, and I still wonder if its the best. Say one want to use the following yaml content in a file mylist.yaml
&#x60;&#x60;&#x60;&#x60;yaml
-firstkey: firstvalue
-secondkey: secondvalue
&#x60;&#x60;&#x60;&#x60;
and you want to loop over the items of the list and access the keys and values. Then, the only solution I could find was to write
&#x60;&#x60;&#x60;&#x60;liquid
{% for thing in site.data.mylist %}
  {% for hash in thing %}
    {{hash[1]}}
    {{hash[0]}}
  {% endfor %}
{% endfor %}
&#x60;&#x60;&#x60;&#x60;
and it took me a while to figure it out (it doesn&#x27;t seem obvious that one needs to nest two for&#x27;s, especially since the second one only ever loops through one iteration).

Another example: to actually use a variable in the description of the file name in the loop
&#x60;&#x60;&#x60;&#x60;liquid
{% for thing in site.data.myfile %}
&#x60;&#x60;&#x60;&#x60;
one cas surround it in brackets, like:
&#x60;&#x60;&#x60;&#x60;liquid
{% for thing in site.data.[page.layout] %}
&#x60;&#x60;&#x60;&#x60;
(which will loop in site.data.foo if page.layout is foo).
Again, it took me a while to figure out this syntax.
Etc.

Maybe I am missing some documentation somewhere, but I did do a good search before posting this suggestion.

Anyway, great job, keep up the good work ;-)


### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73106146) on: **Thursday, February 05, 2015**

Yaml feels very natural to everybody who is familiar with object-oriented languages. It’s very similar to Json. Thats why it is not documented here.

&gt; To make things concrete, here is a problem for which I only have an ugly solution

You show a *solution*. What is the problem you are trying to solve?
---
> from: [**baptistecalmes**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73231451) on: **Friday, February 06, 2015**

I&#x27;m not trying to solve any problem.
I&#x27;m not complaining about documenting yaml.

I&#x27;m just saying that it would help new Jekyll users if the way *liquid* is used to parse yaml was more documented on the jekyll home page.

(and then I gave two examples of things that are not so easy to guess at first sight, from the documentation available)
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73253978) on: **Friday, February 06, 2015**

I&#x27;m also looking for more information about looping through YML values. Suppose I have 30 pages in my site and want to loop through them in a specific order as described in the YML file? It would be great to have a code sample showing how to do that.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73284666) on: **Friday, February 06, 2015**

The Liquid creators have an [introduction video](http://youtu.be/tZLTExLukSg) on top of the [official documentation](http://docs.shopify.com/themes/liquid-documentation/basics).

Try to put this in your &#x60;_config.yml&#x60;:

&#x60;&#x60;&#x60;yaml
numbers:
  - zero
  - one
  - two
&#x60;&#x60;&#x60;

Now you can loop over the &#x60;site.numbers&#x60; array:

&#x60;&#x60;&#x60;twig
{% for number in site.numbers %}
  {{ number | upcase }}
{% endfor %}
&#x60;&#x60;&#x60;

The result is:

&#x60;&#x60;&#x60;
ZERO
ONE
TWO
&#x60;&#x60;&#x60;

The array items are accessible over the index too.

&#x60;&#x60;&#x60;twig
{{ site.numbers[1] | upcase }}

#=&gt; ONE
&#x60;&#x60;&#x60;

----

For more useful examples checkout the [Jekyll Snippets](https://github.com/mdo/jekyll-snippets). You can find a heavier Liquid usage in my [Compress HTML layout](https://github.com/penibelst/jekyll-compress-html).

---
> from: [**baptistecalmes**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73318013) on: **Friday, February 06, 2015**

The official documentation of liquid does not mention yaml at all, and the parsing of yaml files with liquid is a feature that is implemented by Jekyll (or maybe by some ruby package it relies upon). It does not come from the original liquid specs.

Am I right, or am I missing something?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73322646) on: **Friday, February 06, 2015**

You are right.

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73323646) on: **Friday, February 06, 2015**

@baptistecalmes Feel free to expand the documentation in the jekyll/jekyll repository. It&#x27;s open.

---
> from: [**baptistecalmes**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73326847) on: **Friday, February 06, 2015**

I was afraid you&#x27;d say that ;-).

However, before offering some kind of contribution, I was hoping someone would point me to the right place that I had missed.

If no one else reacts, I&#x27;ll think about it.

Thanks anyway for your time and suggestions.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-73341661) on: **Friday, February 06, 2015**

@penibelst Thanks for the concise examples and quality links. Made a note of the links to stay in tune with.
---
> from: [**mvaneijgen**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-74333346) on: **Friday, February 13, 2015**

I have a problem like this, I am new and don&#x27;t know where to get started or how to call what im looking for.

I have a data set in YAML looking something like this (there is 100+ of these sets)

    - name: 
      image: 
      website:   
      dateadd: 
      releasedate: true|false

And I loop though them using 

&#x60;{% for data in site.data.games %}&#x60;

But now I want to create a separate page where only the &#x27;games&#x27; that do have a release date, so 

&#x60;releasedate: true&#x60;

But I don&#x27;t know where to start or how to call what im doing to get a proper Google search term, anyone tips?

[here is the website btw if someone is interested](http://mvaneijgen.nl/upcoming-vita-games)
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-74970532) on: **Wednesday, February 18, 2015**

mvaneijgen - I think you just need a where clause:

see http://jekyllrb.com/docs/templates/ for a little documentation. It should be easy to do.

Also, it looks like you are putting the data set in the config.yml? you can also put it in a data file (http://jekyllrb.com/docs/datafiles/) which I find very handy.

I find StackOverflow to be the best source of examples, that and random blogs. I think their is an infinite number of things that seem obvious to the one trying to implement it, but that creates a bit of a documentation problem.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-75018414) on: **Thursday, February 19, 2015**

@mvaneijgen 

&#x60;&#x60;&#x60;
{% assign released_games = site.data.games | where: &quot;releasedate&quot;, true %}
{% for game in released_games %}
  {{ game.name }}
{% endfor %}

---
> from: [**mvaneijgen**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-75022403) on: **Thursday, February 19, 2015**

Friendly user on Reddit had pointed me to the solution 
&#x60;&#x60;&#x60;
{% for data in site.data.games %}
   {% if games.releasedate == true %} Do stuff {% endif %}
{% endfor %}
&#x60;&#x60;&#x60;
This works for me 
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/266#issuecomment-75102515) on: **Thursday, February 19, 2015**

@mvaneijgen This can not work. The iteration variable is &#x60;data&#x60;, not &#x60;games&#x60;.

&#x60;&#x60;&#x60;
{% for data in site.data.games %}
   {% if data.releasedate == true %} Do stuff {% endif %}
{% endfor %}
&#x60;&#x60;&#x60;

If you know about [truthy and falsy in Liquid](http://docs.shopify.com/themes/liquid-documentation/basics/true-and-false), you can simplify it.

&#x60;&#x60;&#x60;
{% for game in site.data.games %}
   {% if game.releasedate %} Do stuff {% endif %}
{% endfor %}
&#x60;&#x60;&#x60;

