# Automated Go Class Documentation Workflow

This repository automatically generates and deploys documentation for your Go classes using Hugo and GitHub Pages.

## How It Works

1. **Create a new Go file** (e.g., `class48.go`)
2. **Generate documentation** using the script
3. **Push to GitHub** - Documentation website updates automatically!

## Quick Start

### Step 1: Create a New Go Class File

Create your new Go file (e.g., `class48.go`):

```go
package main

import "fmt"

// HelloWorld demonstrates a simple function
func HelloWorld() {
    fmt.Println("Hello, World from Class 48!")
}

func main() {
    HelloWorld()
}
```

### Step 2: Generate Documentation

Run the documentation generator script:

```bash
./generate-class-doc.sh class48.go
```

This will:
- Create documentation in `docs/content/docs/classes/class48.md`
- Update the classes index
- Show you the next steps

### Step 3: Push to GitHub

```bash
git add .
git commit -m "Add class48.go and generate documentation"
git push origin main
```

### Step 4: View Your Documentation

GitHub Actions will automatically:
- Build the Hugo site
- Deploy to GitHub Pages
- Your documentation will be available at: `https://yourusername.github.io/your-repo-name`

## What Gets Generated

For each Go file, the system creates:

1. **Individual class page** - Complete source code and analysis
2. **Classes index** - List of all Go files with descriptions
3. **API documentation** - Function and type listings
4. **Searchable website** - Hugo-powered documentation site

## Manual Documentation Generation

You can also generate documentation locally:

```bash
# Generate for a specific file
./generate-class-doc.sh your-file.go

# Build Hugo site locally (optional)
cd docs
hugo server --buildDrafts
```

## Directory Structure

```
.
├── class47.go                    # Your Go files
├── class48.go                    # New Go files you create
├── generate-class-doc.sh         # Documentation generator script
├── docs/                         # Hugo documentation site
│   ├── hugo.toml                # Hugo configuration
│   ├── content/
│   │   ├── docs/
│   │   │   ├── classes/         # Auto-generated class docs
│   │   │   ├── api/             # API reference
│   │   │   └── project-overview/
│   │   └── menu/
│   └── themes/book/             # Hugo Book theme
└── .github/
    └── workflows/
        └── docs.yml             # GitHub Actions workflow
```

## Features

- ✅ **Automatic deployment** - Push and your site updates
- ✅ **Source code highlighting** - Syntax-highlighted Go code  
- ✅ **Responsive design** - Mobile-friendly documentation
- ✅ **Search functionality** - Find content quickly
- ✅ **Version control** - Track documentation changes
- ✅ **Free hosting** - GitHub Pages at no cost

## GitHub Pages Setup

To enable GitHub Pages for your repository:

1. Go to your repository on GitHub
2. Click **Settings** > **Pages**
3. Under **Source**, select **GitHub Actions**
4. The workflow will handle the rest!

## Customization

### Adding Comments to Your Go Files

Add comments to your Go files for better documentation:

```go
// Package main demonstrates HTTP server concepts
package main

// TProduct represents a product in our system
type TProduct struct {
    ID    int    // Unique identifier
    Title string // Product name
}

// GetProducts returns all available products
func GetProducts() []TProduct {
    // Implementation here
}
```

### Customizing Hugo Theme

Edit `docs/hugo.toml` to customize:
- Site title and description
- Theme colors and fonts
- Menu structure
- Search settings

## Troubleshooting

### Common Issues

1. **Hugo build fails**: Check `docs/hugo.toml` syntax
2. **GitHub Pages not updating**: Verify GitHub Actions permissions
3. **Script permission denied**: Run `chmod +x generate-class-doc.sh`

### Viewing Build Logs

Check the **Actions** tab in your GitHub repository to see build logs and debug any issues.

---

Happy coding! Your Go classes will now have beautiful, automatically-updated documentation! 🚀
