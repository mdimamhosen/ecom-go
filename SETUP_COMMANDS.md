# Quick Setup Commands

## Initialize Git Repository (if not already done)
```bash
git init
git add .
git commit -m "Initial commit with automated documentation workflow"
```

## Create GitHub Repository
1. Go to GitHub.com
2. Create a new repository named "ecom-go" (or your preferred name)
3. Don't initialize with README (since we already have files)

## Connect to GitHub and Push
```bash
git remote add origin https://github.com/yourusername/ecom-go.git
git branch -M main
git push -u origin main
```

## Enable GitHub Pages
1. Go to your repository on GitHub
2. Click Settings > Pages
3. Under Source, select "GitHub Actions"
4. The workflow will automatically deploy your documentation!

## Test the Workflow
```bash
# Create a new Go file
echo 'package main
import "fmt"
func Class49Demo() { fmt.Println("Hello Class 49!") }' > class49.go

# Generate documentation
./generate-class-doc.sh class49.go

# Commit and push
git add .
git commit -m "Add class49.go with auto-generated docs"
git push origin main
```

Your documentation website will be available at:
https://yourusername.github.io/ecom-go/
