#!/bin/bash

# Helper script to suggest the next class number and create a new class file
# Usage: ./new-class.sh [topic-description]

echo "🔍 Finding the next available class number..."

# Find the highest class number
HIGHEST=$(ls class*.go 2>/dev/null | grep -o 'class[0-9]*' | grep -o '[0-9]*' | sort -n | tail -1)

if [ -z "$HIGHEST" ]; then
    NEXT=1
else
    NEXT=$((HIGHEST + 1))
fi

NEXT_FILE="class${NEXT}.go"

echo "📊 Current class files:"
ls class*.go 2>/dev/null | sort -V | tail -5 || echo "   (No class files found)"

echo ""
echo "✨ Next class number: $NEXT"
echo "📝 Next file name: $NEXT_FILE"

if [ $# -gt 0 ]; then
    TOPIC="$*"
    echo "🎯 Topic: $TOPIC"
else
    echo "💡 Example usage: $0 \"JWT Authentication\""
    echo ""
    read -p "📝 Enter topic description (optional): " TOPIC
fi

echo ""
read -p "🚀 Create $NEXT_FILE now? (y/N): " -n 1 -r
echo

if [[ $REPLY =~ ^[Yy]$ ]]; then
    if [ -f "$NEXT_FILE" ]; then
        echo "⚠️  File $NEXT_FILE already exists!"
        exit 1
    fi
    
    # Copy template and customize
    cp class-template.go "$NEXT_FILE"
    
    # Replace placeholders in the new file
    if [ ! -z "$TOPIC" ]; then
        sed -i "s/ClassXX demonstrates \[DESCRIBE WHAT THIS CLASS TEACHES\]/Class${NEXT} demonstrates ${TOPIC}/g" "$NEXT_FILE"
        sed -i "s/Class XX - \[Your Topic\]/Class ${NEXT} - ${TOPIC}/g" "$NEXT_FILE"
        sed -i "s/\[Brief description of what this class demonstrates\]/${TOPIC}/g" "$NEXT_FILE"
        sed -i "s/This class demonstrates: \[Your learning objective\]/This class demonstrates: ${TOPIC}/g" "$NEXT_FILE"
        sed -i "s/Class XX Demo: \[Your Topic\]/Class ${NEXT} Demo: ${TOPIC}/g" "$NEXT_FILE"
    fi
    
    # Replace class numbers
    sed -i "s/ClassXX/Class${NEXT}/g" "$NEXT_FILE"
    sed -i "s/classXX/class${NEXT}/g" "$NEXT_FILE"
    sed -i "s/Class XX/Class ${NEXT}/g" "$NEXT_FILE"
    
    echo "✅ Created $NEXT_FILE"
    echo ""
    echo "📋 Next steps:"
    echo "   1. Edit the file: nano $NEXT_FILE"
    echo "   2. Implement your code"
    echo "   3. Generate docs: ./generate-class-doc.sh $NEXT_FILE"
    echo "   4. Commit: git add . && git commit -m \"Add $NEXT_FILE: $TOPIC\" && git push"
    echo ""
    echo "🎯 File created with template structure. Happy coding!"
else
    echo "❌ Cancelled."
    echo ""
    echo "💡 Manual steps:"
    echo "   cp class-template.go $NEXT_FILE"
    echo "   nano $NEXT_FILE"
    echo "   ./generate-class-doc.sh $NEXT_FILE"
fi
