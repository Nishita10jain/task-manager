const API_URL = "http://localhost:8080/tasks";

// Fetch and display tasks
async function fetchTasks() {
    const response = await fetch(API_URL);
    const tasks = await response.json();

    const tasksList = document.getElementById("tasksList");
    tasksList.innerHTML = "";

    tasks.forEach(task => {
        const taskItem = document.createElement("div");
        taskItem.innerHTML = `
            <p><strong>${task.title}</strong></p>
            <p>${task.description}</p>
            <p>Due: ${task.due_date}</p>
            <p>Status: ${task.done ? "✅ Completed" : "❌ Pending"}</p>
            <button onclick="updateTask(${task.id})">Mark as Done</button>
        `;
        tasksList.appendChild(taskItem);
    });
}

// Handle task creation
document.getElementById("taskForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const title = document.getElementById("title").value;
    const description = document.getElementById("description").value;
    const dueDate = document.getElementById("dueDate").value;

    await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, description, due_date: dueDate, done: false })
    });

    fetchTasks();
});

// Load tasks on page load
fetchTasks();
