<?php
$servername = "localhost";
$username = "root";
$password = "";
$database = "crud";
$con = mysqli_connect($servername, $username, $password, $database);
if ($con->connect_error) {
    die("Connect failed :" . $con->connect_error);
}
