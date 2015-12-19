# [Troubleshooting on OS X 10.9](https://github.com/jekyll/jekyll-help/issues/108)

> state: **closed** opened by: **** on: ****

There are a number of people who are having problems with installing Jekyll on OS X 10.9 based on the number of work-arounds posted elsewhere.

A good example can be found here. http://stackoverflow.com/questions/22479246/how-to-install-jekyll-on-mac-osx-10-9-with-ruby-2-0-0

The troubleshooting guide should be updated to include directions to run in the terminal:

1. ruby -e &quot;$(curl -fsSL https://raw.github.com/Homebrew/homebrew/go/install)&quot;
2. brew doctor
3. (follow any hints that doctor might have come up with - adding yourself to group wheel with &quot;dscl . append /Groups/wheel GroupMembership &lt;your short username&gt;&quot; fixes most of them but you might still have to edit your $PATH in your ~/.profile)
4. curl -L https://get.rvm.io | bash -s stable --autolibs=enabled
5. source ~/.rvm/scripts/rvm
6. gem install jekyll

Not that running gem as sudo isn&#x27;t always such a good thing to do as it messes up permissions for later installations.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/108#issuecomment-55642847) on: **Monday, September 15, 2014**

I&#x27;d like to keep the issues here as specific as possible. If someone is having trouble getting Jekyll up and running on their machine, we shouldn&#x27;t assume anything about what they&#x27;ve done so far.

That said, this is still valuable and we will probably reference it in the future. Thanks!
