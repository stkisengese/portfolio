import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Home.css';

const Home = () => {
  const [about, setAbout] = useState(null);

  useEffect(() => {
    const fetchAbout = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/about');
        setAbout(response.data);
      } catch (error) {
        console.error('Error fetching about data:', error);
      }
    };

    fetchAbout();
  }, []);

  if (!about) return <div>Loading...</div>;

  return (
    <section className="home">
      <div className="hero">
        <img src="/profile_pic.jpeg" alt={about.name} className="profile-image" />
        <h1>{about.name}</h1>
        <h2>{about.role}</h2>
        <p>{about.description}</p>
      </div>
    </section>
  );
};

export default Home;