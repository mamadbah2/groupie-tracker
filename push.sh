# Automatisation des commits
while true; do
    # git config --global credential.helper store
    # git config --global user.email "janelaffranchis@gmail.com"
    # git config --global user.name "Franchis Janel MOKOMBA"
    git add .
    git commit -m "Janel Template"
    git push
    sleep 30 # 30 seconds
done