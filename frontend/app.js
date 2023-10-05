const express = require('express');
const bodyParser = require('body-parser');
const { Pool } = require('pg');


const app = express();

const pool = new Pool({
  user: 'postgres',  // replace with your PostgreSQL username
  host: 'hestia-backend', // replace with the name or ID of your PostgreSQL container
  database: 'hestia', // replace with your PostgreSQL database name
  password: '123', // replace with your PostgreSQL password
  port: 5432, // replace with your PostgreSQL server port number
});

app.use(express.static('public'));
app.use(express.urlencoded({ extended: true }));

app.set('view engine', 'pug');

app.get('/', async (req, res) => {
  try {
    const result = await pool.query('SELECT * FROM Day WHERE Date = CURRENT_DATE');
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

app.post('/addFood', function(req, res) {
  const foodName = req.body.name;
  const query = `INSERT INTO Tag (Name, Date, Food_id)
                 SELECT '${foodName}', NOW(), Food_id
                 FROM Food
                 WHERE Name = '${foodName}'`;
  pool.query(query, (err, result) => {
    if (err) {
      console.error(err);
      res.send('Error inserting food into Tag table');
    } else {
      res.redirect('/table');
    }
  });
});

app.post('/addWeightTraining', function(req, res) {
  const exerciseName = req.body.name;
  const query = `INSERT INTO Day (Name, Date, Weight_id)
                 SELECT '${exerciseName}', NOW(), Weight_id
                 FROM Weight
                 WHERE Name = '${exerciseName}'`;
  pool.query(query, (err, result) => {
    if (err) {
      console.error(err);
      res.send('Error inserting exercise into Tag table');
    } else {
      res.redirect('/table');
    }
  });
});

app.post('/addEnduranceTraining', function(req, res) {
  const exerciseName = req.body.name;
  const query = `INSERT INTO Day (Name, Date, Cardio_id)
                 SELECT '${exerciseName}', NOW(), Cardio_id
                 FROM Cardio
                 WHERE Name = '${exerciseName}'`;
  pool.query(query, (err, result) => {
    if (err) {
      console.error(err);
      res.send('Error inserting exercise into Tag table');
    } else {
      res.redirect('/table');
    }
  });
});


app.listen(3000, () => {
  console.log('Server listening on http://localhost:3000');
});


