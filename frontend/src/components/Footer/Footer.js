import React from 'react';
import './Footer.css';

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <div className="social-links">
          <a href="https://github.com/stkisengese" target="_blank" rel="noopener noreferrer">GitHub</a>
          <a href="https://linkedin.com/in/stephen-kisengese" target="_blank" rel="noopener noreferrer">LinkedIn</a>
          <a href="https://twitter.com/skisengese" target="_blank" rel="noopener noreferrer">Twitter</a>
        </div>
        <p>&copy; {new Date().getFullYear()} Stephen Kisengese. All rights reserved.</p>
      </div>
    </footer>
  );
};

export default Footer;