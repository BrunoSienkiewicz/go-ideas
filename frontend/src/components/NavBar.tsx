import React from "react";

const NavBar = () => {
  return (
    <div className="u-clearfix u-sheet u-valign-middle u-sheet-1">
      <nav className="u-menu u-menu-one-level u-offcanvas u-menu-1">
        <div className="menu-collapse">
          <a className="u-button-style u-nav-link" href="#">
            <svg className="u-svg-link" viewBox="0 0 24 24"></svg>
            <svg
              className="u-svg-content"
              version="1.1"
              id="svg-a781"
              viewBox="0 0 16 16"
              x="0px"
              y="0px"
              xmlns="http://www.w3.org/2000/svg"
            >
              <g>
                <rect y="1" width="16" height="2"></rect>
                <rect y="7" width="16" height="2"></rect>
                <rect y="13" width="16" height="2"></rect>
              </g>
            </svg>
          </a>
        </div>
        <div className="u-custom-menu u-nav-container">
          <ul className="u-nav u-unstyled">
            <li className="u-nav-item">
              <a className="u-button-style u-nav-link" href="Główna.html">
                Główna
              </a>
            </li>
            <li className="u-nav-item">
              <a className="u-button-style u-nav-link" href="Wyświetl.html">
                Wyświetl
              </a>
            </li>
            <li className="u-nav-item">
              <a className="u-button-style u-nav-link" href="Dodaj.html">
                Dodaj
              </a>
            </li>
          </ul>
        </div>
        <div className="u-custom-menu u-nav-container-collapse">
          <div className="u-black u-container-style u-inner-container-layout u-opacity u-opacity-95 u-sidenav">
            <div className="u-inner-container-layout u-sidenav-overflow">
              <div className="u-menu-close"></div>
              <ul className="u-align-center u-nav u-popupmenu-items u-unstyled u-nav-2">
                <li className="u-nav-item">
                  <a className="u-button-style u-nav-link" href="Główna.html">
                    Główna
                  </a>
                </li>
                <li className="u-nav-item">
                  <a className="u-button-style u-nav-link" href="Wyświetl.html">
                    Wyświetl
                  </a>
                </li>
                <li className="u-nav-item">
                  <a className="u-button-style u-nav-link" href="Dodaj.html">
                    Dodaj
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="u-black u-menu-overlay u-opacity u-opacity-70"></div>
        </div>
      </nav>
      <nav className="u-menu u-menu-one-level u-offcanvas u-menu-2">
        <div className="u-custom-menu u-nav-container">
          <ul className="u-nav u-unstyled u-nav-3">
            <li className="u-nav-item">
              <a
                className="u-button-style u-nav-link u-text-active-palette-1-base u-text-hover-palette-2-base"
                href="Logowanie.html"
                style={{ padding: "10px 20px" }}
              >
                Logowanie
              </a>
            </li>
            <li className="u-nav-item">
              <a
                className="u-button-style u-nav-link u-text-active-palette-1-base u-text-hover-palette-2-base"
                href="Rejestracja.html"
                style={{ padding: "10px 20px" }}
              >
                Rejestracja
              </a>
            </li>
          </ul>
        </div>
        <div className="u-custom-menu u-nav-container-collapse">
          <div className="u-black u-container-style u-inner-container-layout u-opacity u-opacity-95 u-sidenav">
            <div className="u-inner-container-layout u-sidenav-overflow">
              <div className="u-menu-close"></div>
              <ul className="u-align-center u-nav u-popupmenu-items u-unstyled u-nav-4">
                <li className="u-nav-item">
                  <a
                    className="u-button-style u-nav-link"
                    href="Logowanie.html"
                  >
                    Logowanie
                  </a>
                </li>
                <li className="u-nav-item">
                  <a
                    className="u-button-style u-nav-link"
                    href="Rejestracja.html"
                  >
                    Rejestracja
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="u-black u-menu-overlay u-opacity u-opacity-70"></div>
        </div>
      </nav>
    </div>
  );
};

export default NavBar;
