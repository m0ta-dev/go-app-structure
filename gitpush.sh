eval "$(ssh-agent -s)"
ssh-add ~/.ssh/gitmota
git push -f github main
git push -f gitlab main