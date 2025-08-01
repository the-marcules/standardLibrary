"use client";
import React, { useEffect } from "react";
import { NextPage } from "next";

const NotFound: NextPage = () => {
  useEffect(() => {
    window.location.href = "/de/404"; // Redirect to home page
  }, []);

  return <></>;
};

export default NotFound;
