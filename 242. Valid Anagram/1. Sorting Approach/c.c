#include <stdio.h>
#include <string.h>
#include <stdbool.h>

// Helper function to sort a string
void sortString(char *str) {
    int n = strlen(str);
    for (int i = 0; i < n - 1; i++) {
        for (int j = i + 1; j < n; j++) {
            if (str[i] > str[j]) {
                // Swap characters if they are in the wrong order
                char temp = str[i];
                str[i] = str[j];
                str[j] = temp;
            }
        }
    }
}

bool isAnagram(char *s, char *t) {
    // If lengths don't match, they can't be anagrams
    if (strlen(s) != strlen(t)) {
        return false;
    }
    
    // Sort both strings
    sortString(s);
    sortString(t);
    
    // Compare sorted strings
    for (int i = 0; s[i] != '\0'; i++) {
        if (s[i] != t[i]) {
            return false;
        }
    }
    
    return true;
}

int main() {
    char s1[] = "anagram";
    char t1[] = "nagaram";
    char s2[] = "rat";
    char t2[] = "car";
    
    printf("Is Anagram (s1, t1): %s\n", isAnagram(s1, t1) ? "true" : "false");
    printf("Is Anagram (s2, t2): %s\n", isAnagram(s2, t2) ? "true" : "false");
    
    return 0;
}
