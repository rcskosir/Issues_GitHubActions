import re
import sys

def update_changelog(input_string, changelog_file='CHANGELOG.md'):
    # Step 1: Identify the type of entry and clean the input string
    match = re.match(r'^\[(BUG|ENHANCEMENT|FEATURE|BREAKING)\](.*)', input_string.strip())
    if not match:
        print("Error: Input string must start with '[BUG]' or '[ENHANCEMENT]' or '[FEATURE]' or '[BREAKING]'.")
        return
    
    change_type = match.group(1)
    cleaned_string = match.group(2).strip()

    # Step 2: Determine the appropriate header
    if change_type == 'BUG':
        header = 'BUG FIXES:'
    elif change_type == 'ENHANCEMENT':
        header = 'ENHANCEMENTS:'
    elif change_type == 'FEATURE':
        header = 'FEATURES:'
    elif change_type == 'BREAKING':
      header = 'BREAKING CHANGES:'
        
    # Step 3: Read the existing changelog file and look for the header
    with open(changelog_file, 'r') as file:
        lines = file.readlines()

    # Step 4: Find the first occurrence of the header and insert the string under it
    header_found = False
    new_lines = []

    for line in lines:
        new_lines.append(line)
        
        # If the header is found, insert the cleaned string after it
        if header_found == False and header in line:
            new_lines.append(f"{cleaned_string}\n")
            header_found = True

    # If the header is not found, append the entry at the end of the file
    if not header_found:
        print(f"Warning: '{header}' not found in the changelog. Appending the entry at the top.")
        new_lines.insert(0, f"\n### {header}\n")
        new_lines.insert(1, f"{cleaned_string}\n")

    # Step 5: Write the updated content back to the file
    with open(changelog_file, 'w') as file:
        file.writelines(new_lines)

    print(f"The change has been added to the changelog under {header}.")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Please provide the input string as a command-line argument.")
        sys.exit(1)

    input_string = sys.argv[1]  # Get the entry from the command line
    
    update_changelog(input_string)
