function copyToClipboard() {
    const flashLink = document.getElementById("shortLink");
    if (flashLink) {
        // Create a temporary input to copy the text
        const tempInput = document.createElement("input");
        tempInput.value = flashLink.textContent;
        document.body.appendChild(tempInput);

        // Select and copy the content
        tempInput.select();
        tempInput.setSelectionRange(0, 99999); // For mobile devices
        document.execCommand("copy");

        // Remove the temporary input
        document.body.removeChild(tempInput);

        // Provide feedback to the user
    }
}
