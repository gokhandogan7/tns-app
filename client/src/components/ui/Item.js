/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";

const Item = ({ items }) => {
  return items
    
    .map((item, index) => (
      <div data-cy="article" className="container" key={index}>
        <div className="itemContainer">
          <p data-cy="title" className="title">{item.Title}</p>
          <p data-cy="description" className="description">{item.desc}</p>
          <p data-cy="content" className="content">{item.context}</p>
        </div>
      </div>
    ));
};

export default Item;
