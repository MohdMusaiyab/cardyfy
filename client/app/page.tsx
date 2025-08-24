import React from "react";
import Header from "./components/Header";
import UserNameInput from "./components/UserNameInput";
import Footer from "./components/Footer";

const page = () => {
  return (
    <div>
      <Header />
      <UserNameInput />
      <Footer />
    </div>
  );
};

export default page;
