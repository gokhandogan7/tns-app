/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const Item = ({ items }) => {
  const navigate = useNavigate();

  const navigateToAuthors = (val) => {
    // 👇️ navigate to /contacts
    navigate(`/fullarticles/${val}`);
  };

  return items
    
    .map((item, index) => (
      <div onClick={()=>navigateToAuthors(item.Id)} data-cy="article" className="container" key={index}>
        <div className="itemContainer">
          <p data-cy="title" className="title">{item.Title}</p>
          <p data-cy="description" className="description">{item.desc}</p>
          <p data-cy="content" className="content">{item.Text}</p>
          <p data-cy="email" className="content">{item.Short_Text}</p>
          <p data-cy="user" className="content">{item.Name}</p>
          <p data-cy="email" className="content">{item.Email}</p>
          <p data-cy="email" className="content">{item.Date}</p>
          
        </div>
      </div>
    ));
};

export default Item;
