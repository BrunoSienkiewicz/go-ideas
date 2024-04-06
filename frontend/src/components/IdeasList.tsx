import React, { useState, useEffect } from "react";

interface Idea {
  name: string;
  category: string;
  attributes: Attribute[];
}

interface Attribute {
  name: string;
  value: string;
}

const IdeasList: React.FC = () => {
  const [ideas, setIdeas] = useState<Idea[]>([]);

  useEffect(() => {
    fetchIdeas();
  }, []);

  const fetchIdeas = async () => {
    try {
      const response = await fetch("http://localhost:5000/idea");
      if (!response.ok) {
        throw new Error("Failed to fetch ideas");
      }
      const ideasData = await response.json();
      setIdeas(ideasData);
    } catch (error) {
      console.error("Error fetching ideas:", error);
    }
  };

  return (
    <div>
      <h1>Ideas List</h1>
      {ideas.map((idea, index) => (
        <div key={index}>
          <h2>{idea.name}</h2>
          <p>Category: {idea.category}</p>
          <h3>Attributes:</h3>
          <ul>
            {idea.attributes.map((attribute, attributeIndex) => (
              <li key={attributeIndex}>
                <strong>{attribute.name}: </strong>
                {attribute.value}
              </li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
};

export default IdeasList;
