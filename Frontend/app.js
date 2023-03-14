const express = require('express');
const { Pool } = require('pg');

const app = express();

const pool = new Pool({
  user: 'postgres',  // replace with your PostgreSQL username
  host: 'localhost', // replace with your PostgreSQL server hostname or IP address
  database: 'hestia', // replace with your PostgreSQL database name
  password: '123', // replace with your PostgreSQL password
  port: 5432, // replace with your PostgreSQL server port number
});

app.use(express.static('public'));

app.set('view engine', 'pug');

app.get('/', async (req, res) => {
  try {
    const result = await pool.query('SELECT * FROM essen');
    const rows = result.rows;
    res.render('index', { rows });
  } catch (err) {
    console.error(err);
    res.status(500).send('Internal Server Error');
  }
});

app.get('/table', async (req, res) => {
  res.render('table');
});

app.get('/dashboard', async (req, res) => {
    res.render('dashboard');
});

app.get('/food', async (req, res) => {
  res.render('food');
});

app.listen(3000, () => {
  console.log('Server listening on http://localhost:3000');
});

