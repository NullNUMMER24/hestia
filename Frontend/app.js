const { Pool } = require('pg');
const express = require('express');
const ejs = require('ejs');

const pool = new Pool({
  user: 'postgres',  // replace with your PostgreSQL username
  host: 'localhost', // replace with your PostgreSQL server hostname or IP address
  database: 'hestia', // replace with your PostgreSQL database name
  password: '123', // replace with your PostgreSQL password
  port: 5432, // replace with your PostgreSQL server port number
});

const app = express();

app.set('view engine', 'ejs');

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

app.listen(3000, () => {
  console.log('Server listening on http://localhost:3000');
});

