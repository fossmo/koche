✨ Koche is a simple tool that helps software development teams ensure that commit messages
adhere to the Conventional Commits standard. 

### What are Conventional Commits?

Conventional Commits are a standardized naming convention for Git commit messages that 
provide consistency and clarity in version control histories.

Conventional Commits consists of a structured message that includes a type (e.g., feat, fix, chore), 
an optional scope in parentheses, an optional exclamation mark for breaking changes, 
and a description. Read more at https://www.conventionalcommits.org

## 💡 Usage with Git 

From your git repository run the following command:

```
    koche -i
```

This creates a git hook that ensure that commit messages
adhere to the Conventional Commits standard.