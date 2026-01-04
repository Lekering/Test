const API_BASE_URL = '/api';

// Состояние приложения
let currentFilter = 'all'; // 'all' или 'active'
let tasks = {};

// Элементы DOM
const taskForm = document.getElementById('taskForm');
const taskTitle = document.getElementById('taskTitle');
const taskDescription = document.getElementById('taskDescription');
const tasksContainer = document.getElementById('tasksContainer');
const showAllBtn = document.getElementById('showAll');
const showActiveBtn = document.getElementById('showActive');
const loadingEl = document.getElementById('loading');

// Инициализация
document.addEventListener('DOMContentLoaded', () => {
    loadTasks();
    
    taskForm.addEventListener('submit', handleCreateTask);
    showAllBtn.addEventListener('click', () => setFilter('all'));
    showActiveBtn.addEventListener('click', () => setFilter('active'));
});

// Загрузка задач
async function loadTasks() {
    try {
        loadingEl.style.display = 'block';
        tasksContainer.innerHTML = '';
        
        const url = currentFilter === 'active' 
            ? `${API_BASE_URL}/tasks?done=false`
            : `${API_BASE_URL}/tasks`;
            
        const response = await fetch(url);
        
        if (!response.ok) {
            throw new Error(`Ошибка загрузки: ${response.status}`);
        }
        
        tasks = await response.json();
        renderTasks();
        
        loadingEl.style.display = 'none';
    } catch (error) {
        loadingEl.style.display = 'none';
        showError(`Не удалось загрузить задачи: ${error.message}`);
    }
}

// Создание задачи
async function handleCreateTask(e) {
    e.preventDefault();
    
    const title = taskTitle.value.trim();
    const description = taskDescription.value.trim();
    
    if (!title || !description) {
        showError('Пожалуйста, заполните все поля');
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, description }),
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.Massege || 'Ошибка создания задачи');
        }
        
        // Очистка формы
        taskForm.reset();
        
        // Перезагрузка задач
        await loadTasks();
        
    } catch (error) {
        showError(`Не удалось создать задачу: ${error.message}`);
    }
}

// Переключение статуса задачи
async function toggleTaskStatus(title, currentStatus) {
    try {
        const response = await fetch(`${API_BASE_URL}/tasks/${encodeURIComponent(title)}`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ done: !currentStatus }),
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.Massege || 'Ошибка обновления задачи');
        }
        
        await loadTasks();
        
    } catch (error) {
        showError(`Не удалось обновить задачу: ${error.message}`);
    }
}

// Удаление задачи
async function deleteTask(title) {
    if (!confirm(`Вы уверены, что хотите удалить задачу "${title}"?`)) {
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE_URL}/tasks/${encodeURIComponent(title)}`, {
            method: 'DELETE',
        });
        
        if (!response.ok) {
            throw new Error('Ошибка удаления задачи');
        }
        
        await loadTasks();
        
    } catch (error) {
        showError(`Не удалось удалить задачу: ${error.message}`);
    }
}

// Отображение задач
function renderTasks() {
    const tasksArray = Object.values(tasks);
    
    if (tasksArray.length === 0) {
        tasksContainer.innerHTML = `
            <div class="empty-state">
                <div class="empty-state-icon">📝</div>
                <div class="empty-state-text">Нет задач</div>
                <div>Создайте свою первую задачу выше</div>
            </div>
        `;
        return;
    }
    
    tasksContainer.innerHTML = tasksArray
        .map(task => createTaskCard(task))
        .join('');
    
    // Добавляем обработчики событий
    tasksArray.forEach(task => {
        const title = task.Title;
        const toggleBtn = document.getElementById(`toggle-${encodeURIComponent(title)}`);
        const deleteBtn = document.getElementById(`delete-${encodeURIComponent(title)}`);
        
        if (toggleBtn) {
            toggleBtn.addEventListener('click', () => toggleTaskStatus(title, task.IsDone));
        }
        
        if (deleteBtn) {
            deleteBtn.addEventListener('click', () => deleteTask(title));
        }
    });
}

// Создание карточки задачи
function createTaskCard(task) {
    const statusClass = task.IsDone ? 'completed' : '';
    const statusBadge = task.IsDone 
        ? '<span class="status-badge completed">✓ Выполнено</span>'
        : '<span class="status-badge active">В работе</span>';
    
    const createdDate = new Date(task.CreateAt).toLocaleString('ru-RU', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
    
    const doneDate = task.DoneAt 
        ? new Date(task.DoneAt).toLocaleString('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
          })
        : '';
    
    return `
        <div class="task-card ${statusClass}">
            <div class="task-header">
                <div style="flex: 1;">
                    <div class="task-title">${escapeHtml(task.Title)}</div>
                    ${statusBadge}
                </div>
            </div>
            <div class="task-description">${escapeHtml(task.Description)}</div>
            <div class="task-meta">
                📅 Создано: ${createdDate}
                ${doneDate ? `<br>✅ Выполнено: ${doneDate}` : ''}
            </div>
            <div class="task-actions">
                <button 
                    id="toggle-${encodeURIComponent(task.Title)}"
                    class="btn ${task.IsDone ? 'btn-secondary' : 'btn-success'}"
                >
                    ${task.IsDone ? '↩️ Отменить выполнение' : '✓ Отметить выполненной'}
                </button>
                <button 
                    id="delete-${encodeURIComponent(task.Title)}"
                    class="btn btn-danger"
                >
                    🗑️ Удалить
                </button>
            </div>
        </div>
    `;
}

// Установка фильтра
function setFilter(filter) {
    currentFilter = filter;
    
    if (filter === 'all') {
        showAllBtn.classList.add('active');
        showActiveBtn.classList.remove('active');
    } else {
        showActiveBtn.classList.add('active');
        showAllBtn.classList.remove('active');
    }
    
    loadTasks();
}

// Отображение ошибки
function showError(message) {
    const errorDiv = document.createElement('div');
    errorDiv.className = 'error-message';
    errorDiv.textContent = message;
    
    tasksContainer.insertBefore(errorDiv, tasksContainer.firstChild);
    
    setTimeout(() => {
        errorDiv.remove();
    }, 5000);
}

// Экранирование HTML
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

