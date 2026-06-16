// Conceptual blueprint for a Phase 4 live autocomplete dropdown
inputField.addEventListener('input', (e) => {
    let query = e.target.value;
    if (query.length < 2) return;

    fetch(`/api/suggestions?q=${query}`)
        .then(res => res.body.json())
        .then(matches => {
            // Update an absolute-positioned dropdown container under the nav bar
        });
});