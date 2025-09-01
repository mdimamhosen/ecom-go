#!/bin/bash

# Script to generate documentation for a new Go class file
# Usage: ./generate-class-doc.sh <classfile.go>

if [ $# -eq 0 ]; then
    echo "❌ Usage: $0 <classfile.go>"
    echo ""
    echo "📝 File Naming Convention:"
    echo "   ✅ class47.go, class48.go, class49.go..."
    echo "   ❌ Class47.go, my-class.go, myclass.go"
    echo ""
    echo "📋 Examples:"
    echo "   $0 class49.go     # Creates documentation for class49.go"
    echo "   $0 class50.go     # Creates documentation for class50.go"
    echo ""
    echo "💡 Tip: Use sequential numbers starting from your last class"
    echo "   Current files: $(ls class*.go 2>/dev/null | tail -3 | xargs)"
    echo ""
    echo "🎯 To create a new class:"
    echo "   1. Copy class-template.go to classXX.go"
    echo "   2. Edit the new file with your code"
    echo "   3. Run: $0 classXX.go"
    exit 1
fi

CLASSFILE=$1
BASENAME=$(basename "$CLASSFILE" .go)

# Validate file name format
if [[ ! "$CLASSFILE" =~ ^class[0-9]+\.go$ ]]; then
    echo "⚠️  Warning: File name '$CLASSFILE' doesn't follow the recommended pattern"
    echo "📝 Recommended format: classXX.go (where XX is a number)"
    echo "   Examples: class47.go, class48.go, class49.go"
    echo ""
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Cancelled. Please rename your file to follow the pattern."
        exit 1
    fi
fi

if [ ! -f "$CLASSFILE" ]; then
    echo "❌ Error: File $CLASSFILE does not exist!"
    echo ""
    echo "💡 To create a new class file:"
    echo "   cp class-template.go $CLASSFILE"
    echo "   nano $CLASSFILE  # Edit with your code"
    echo "   $0 $CLASSFILE    # Generate documentation"
    exit 1
fi

echo "Generating documentation for $CLASSFILE..."

# Create docs directory structure if it doesn't exist
mkdir -p docs/content/docs/classes

# Generate individual class documentation
cat > "docs/content/docs/classes/${BASENAME}.md" << EOF
---
title: "$BASENAME"
weight: 100
date: $(date '+%Y-%m-%dT%H:%M:%S%z')
---

# $BASENAME Documentation

## Source Code

\`\`\`go
$(cat "$CLASSFILE")
\`\`\`

## File Information

- **Filename**: $CLASSFILE
- **Created**: $(date)
- **Size**: $(wc -c < "$CLASSFILE") bytes
- **Lines**: $(wc -l < "$CLASSFILE") lines

## Functions and Types

$(grep -n "^func\|^type" "$CLASSFILE" | sed 's/^/- Line /')

## Description

$(head -n 20 "$CLASSFILE" | grep -E "^//|^/\*" | sed 's|^//||g' | sed 's|^/\*||g' | xargs || echo "No description found in comments.")

---

*This documentation was auto-generated from the source file.*
EOF

echo "Documentation generated: docs/content/docs/classes/${BASENAME}.md"

# Update the classes index
echo "Updating classes index..."

cat > docs/content/docs/classes/_index.md << 'EOF'
---
title: "Go Classes"
weight: 25
bookCollapseSection: false
---

# Go Classes and Files

This section contains documentation for all Go class files in this project.

## Available Classes

EOF

# Add all Go files to the index
for gofile in *.go; do
    if [ -f "$gofile" ]; then
        basename_file=$(basename "$gofile" .go)
        echo "- [$gofile](${basename_file}) - $(head -n 10 "$gofile" | grep -E '^//|^/\*' | head -n 1 | sed 's|^//||' | sed 's|^/\*||' | xargs || echo 'Go source file')" >> docs/content/docs/classes/_index.md
    fi
done

echo "" >> docs/content/docs/classes/_index.md
echo "---" >> docs/content/docs/classes/_index.md
echo "*Documentation auto-generated on $(date)*" >> docs/content/docs/classes/_index.md

echo "Classes index updated: docs/content/docs/classes/_index.md"
echo "Done! You can now commit and push your changes."
echo ""
echo "Next steps:"
echo "1. git add ."
echo "2. git commit -m 'Add $CLASSFILE and generate documentation'"
echo "3. git push origin main"
echo ""
echo "GitHub will automatically generate and deploy the webpage!"
