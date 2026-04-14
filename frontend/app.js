const API_BASE = 'http://127.0.0.1:3000/api';
let token = localStorage.getItem('token') || '';

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    initForms();
    updateAuthStatus();
});

function initTabs() {
    const tabs = document.querySelectorAll('.tab-btn');
    tabs.forEach(tab => {
        tab.addEventListener('click', () => {
            document.querySelectorAll('.tab-btn').forEach(t => t.classList.remove('active'));
            document.querySelectorAll('.tab-content').forEach(c => c.classList.remove('active'));
            tab.classList.add('active');
            document.getElementById(tab.dataset.tab).classList.add('active');
        });
    });
}

function initForms() {
    document.getElementById('registerForm').addEventListener('submit', register);
    document.getElementById('loginForm').addEventListener('submit', login);
    document.getElementById('exchangeForm').addEventListener('submit', createExchangeRate);
    document.getElementById('articleForm').addEventListener('submit', createArticle);
    document.getElementById('logoutBtn').addEventListener('click', logout);
}

function showResponse(text) {
    const box = document.getElementById('responseBox');
    document.getElementById('responseText').textContent = typeof text === 'object' ? JSON.stringify(text, null, 2) : text;
    box.classList.add('show');
}

function closeResponse() {
    document.getElementById('responseBox').classList.remove('show');
}

async function request(url, options = {}) {
    const headers = {
        'Content-Type': 'application/json',
        ...options.headers
    };
    if (token) {
        headers['Authorization'] = token;
    }
    try {
        const response = await fetch(`${API_BASE}${url}`, {
            ...options,
            headers
        });
        const data = await response.json();
        return { ok: response.ok, data, status: response.status };
    } catch (error) {
        showResponse('请求失败: ' + error.message);
        return { ok: false, data: { error: error.message } };
    }
}

async function register(e) {
    e.preventDefault();
    const username = document.getElementById('regUsername').value;
    const password = document.getElementById('regPassword').value;
    
    const res = await request('/auth/register', {
        method: 'POST',
        body: JSON.stringify({ username, password })
    });
    
    showResponse(res.data);
    if (res.ok) {
        document.getElementById('registerForm').reset();
    }
}

async function login(e) {
    e.preventDefault();
    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;
    
    const res = await request('/auth/login', {
        method: 'POST',
        body: JSON.stringify({ username, password })
    });
    
    showResponse(res.data);
    if (res.ok && res.data.token) {
        token = res.data.token;
        localStorage.setItem('token', token);
        updateAuthStatus();
        document.getElementById('loginForm').reset();
    }
}

function logout() {
    token = '';
    localStorage.removeItem('token');
    updateAuthStatus();
}

function updateAuthStatus() {
    const userInfo = document.getElementById('userInfo');
    const logoutBtn = document.getElementById('logoutBtn');
    
    if (token) {
        userInfo.textContent = '✅ 已登录';
        logoutBtn.style.display = 'inline-block';
    } else {
        userInfo.textContent = '未登录';
        logoutBtn.style.display = 'none';
    }
}

async function getExchangeRates() {
    const res = await request('/exchangeRates', { method: 'GET' });
    showResponse(res.data);
}

async function createExchangeRate(e) {
    e.preventDefault();
    const baseCurrency = document.getElementById('baseCurrency').value;
    const targetCurrency = document.getElementById('targetCurrency').value;
    const rate = parseFloat(document.getElementById('rate').value);
    
    const res = await request('/exchangeRates', {
        method: 'POST',
        body: JSON.stringify({ base_currency: baseCurrency, target_currency: targetCurrency, rate })
    });
    
    showResponse(res.data);
    if (res.ok) {
        document.getElementById('exchangeForm').reset();
    }
}

async function getArticles() {
    const res = await request('/articles', { method: 'GET' });
    showResponse(res.data);
}

async function createArticle(e) {
    e.preventDefault();
    const title = document.getElementById('articleTitle').value;
    const content = document.getElementById('articleContent').value;
    const preview = document.getElementById('articlePreview').value;
    
    const res = await request('/articles', {
        method: 'POST',
        body: JSON.stringify({ title, content, preview, likes: 0 })
    });
    
    showResponse(res.data);
    if (res.ok) {
        document.getElementById('articleForm').reset();
    }
}

async function getArticle() {
    const id = document.getElementById('articleId').value;
    if (!id) return;
    
    const res = await request(`/articles/${id}`, { method: 'GET' });
    showResponse(res.data);
}

async function likeArticle() {
    const id = document.getElementById('likeArticleId').value;
    if (!id) return;
    
    const res = await request(`/articles/${id}/like`, { method: 'POST' });
    showResponse(res.data);
}

async function getLikes() {
    const id = document.getElementById('likeArticleId').value;
    if (!id) return;
    
    const res = await request(`/articles/${id}/like`, { method: 'GET' });
    showResponse(res.data);
}