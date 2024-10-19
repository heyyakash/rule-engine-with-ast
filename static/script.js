let latestRuleId = null;

window.onload = function () {
    fetchRules();
};

function fetchRules() {
    fetch('/all')
        .then(response => response.json())
        .then(data => {
            const rulesList = document.getElementById('rulesList');
            const ruleSelect = document.getElementById('ruleSelect');
            rulesList.innerHTML = '';
            ruleSelect.innerHTML = '';

            data.forEach(rule => {
                const ruleBox = document.createElement('div');
                ruleBox.className = 'rule-container';

                const ruleHeader = document.createElement('div');
                ruleHeader.className = 'rule-header';
                ruleHeader.textContent = `Rule: ${rule.rule}`;
                ruleBox.appendChild(ruleHeader);

                const jsonInput = document.createElement('textarea');
                jsonInput.placeholder = 'Enter JSON data. Eg {"age":31, "gender":"male"}';
                ruleBox.appendChild(jsonInput);

                const testButton = document.createElement('button');
                testButton.textContent = 'Test Rule';
                testButton.onclick = function () {
                    testRule(rule._id, jsonInput.value, ruleBox);
                };
                ruleBox.appendChild(testButton);

                const resultDiv = document.createElement('div');
                resultDiv.className = 'rule-result';
                ruleBox.appendChild(resultDiv);

                rulesList.appendChild(ruleBox);

                const option = document.createElement('option');
                option.value = rule._id;
                option.textContent = rule.rule;
                ruleSelect.appendChild(option);
            });
        })
        .catch(error => console.error('Error fetching rules:', error));
}


async function addRule() {
    const ruleInput = document.getElementById('ruleInput');
    const rule = ruleInput.value;

    if (!rule) {
        alert('Please enter a rule string.');
        return;
    }

    const res = await fetch('/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ rule: rule })
    })
    const result = await res.text()
    if (res.ok) {
        alert('Rule added successfully!');
        latestRuleId = ruleInput.value;
        ruleInput.value = '';
        document.getElementById('combineSection').style.display = 'block';
        fetchRules();
    } else {
        alert(result)
    }
}

function testRule(ruleId, jsonData, ruleBox) {
    let json;
    try {
        json = JSON.parse(jsonData);
    } catch (e) {
        alert('Invalid JSON data');
        return;
    }

    fetch('/evaluate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ rule_id: ruleId, data: json })
    })
        .then(response => response.json())
        .then(result => {
            const resultDiv = ruleBox.querySelector('.rule-result');
            if(!result.result) resultDiv.style.color = "red"
            if(result.result) resultDiv.style.color = "green"
            resultDiv.textContent = `Result: ${result.result}`;
        })
        .catch(error => console.error('Error evaluating rule:', error));
}

function cancelCombination(){
    document.getElementById('combineSection').style.display = 'none';
    latestRuleId = ""
}


async function combineWithNewRule() {
    const ruleSelect = document.getElementById('ruleSelect');
    const selectedRules = Array.from(ruleSelect.selectedOptions).map(option => option.textContent);

    if (!latestRuleId) {
        alert('No new rule to combine. Please add a rule first.');
        return;
    }

    if (selectedRules.length === 0) {
        alert('Please select at least one rule to combine with the newly added rule.');
        return;
    }

    selectedRules.push(latestRuleId);

    const res = await fetch('/combine', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ rules: selectedRules })
    })
    const r = await res.text()
    if (res.ok) {
        alert('Rules combined successfully!');
        fetchRules();
        document.getElementById('combineSection').style.display = 'none';
    } else {
        alert('Error combining rules.');
    }

}
