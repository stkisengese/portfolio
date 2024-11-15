import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Skills.css';

const Skills = () => {
  const [skills, setSkills] = useState([]);

  useEffect(() => {
    const fetchSkills = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/skills');
        setSkills(response.data);
      } catch (error) {
        console.error('Error fetching skills:', error);
      }
    };

    fetchSkills();
  }, []);

  return (
    <section className="skills">
      <h2>My Skills</h2>
      <div className="skills-grid">
        {skills.map(skill => (
          <div key={skill.id} className="skill-card">
            <h3>{skill.name}</h3>
            <div className="skill-level">
              <div 
                className="skill-progress" 
                style={{ width: `${skill.level}%` }}
              ></div>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default Skills;