<?php
include("connect.php");

if (isset($_GET['id'])) {
    $id = $_GET['id'];
    $sql = "DELETE FROM users WHERE id = $id";
    $res = $con->query($sql);
    if (!$res) {
        die("Couldn't delete user from database");
        header('location: /user/index.php');
        exit;
    }
    header('location: /user/index.php');
    exit;
}
