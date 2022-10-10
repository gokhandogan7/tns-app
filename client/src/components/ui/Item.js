/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";

const Item = ({ items, search }) => {
  return items
    .filter((item) => {
      if (search === "") {
        return item;
      } else if (
        item?.Title?.replaceAll(" ", "")
          .toLowerCase()
          .includes(search?.toLowerCase()) ||
        item?.desc?.toLowerCase().includes(search?.toLowerCase())
      ) {
        return item;
      }
    })
    .map((item, index) => (
      <div className="container" key={index}>
        <div className="itemContainer">
          <p className="title">{item.Title}</p>
          <p className="description">{item.desc}</p>
          <p className="content">{item.content}</p>
        </div>
      </div>
    ));
};

export default Item;
