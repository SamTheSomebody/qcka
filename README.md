<a id="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Unlicense License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/samthesomebody/qcka">
    <img src="images/logo.png" alt="Logo" width="600" height="200">
  </a>

  <h3 align="center">qcka</h3>

  <p align="center">
    Quick alias is a bash tool to add and manage aliases without having to edit your .bashrc or .bash_aliases.
    <br />
    <br />
    <a href="https://github.com/SamTheSomebody/qcka">View Demo <i>(TODO)</i></a>
    &middot;
    <a href="https://github.com/SamTheSomebody/qcka/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    &middot;
    <a href="https://github.com/SamTheSomebody/qcka/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Quick alias allows you to add, manage, remove and view your bash aliases on the fly. It has simple grouping to help keep things organized. 

Here's why:
* Don't stop what you're doing. You're using aliases to speed up your workflow, so speed up managing them! 
* A 'here' function to bookmark your current directory.
* Alias validation *(TODO)* and override dectection for inbuilts, commands, funcitons and other aliases.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

<div align="center">
<img src="https://go.dev/images/go-logo-white.svg" alt="Go" width="100">
&nbsp;
&nbsp;
&nbsp;
&nbsp;
<img src="https://github.com/user-attachments/assets/cbc3adf8-0dff-46e9-a88d-5e2d971c169e" alt="Go" width="100">
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

I'm still working on a release of qcka, for now it can only be built from source.

### Prerequisites

* Go: https://go.dev/doc/install

### Installation

To build from source:
1) Clone this repository.
2) Change to it's directory.
3) Run the following command:
```
go build && export PATH=$PATH:$(go list -f '{{.Target}}') && go install
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Run quick alias with the command: ```qcka```

Quick alias features a built-in manual, alongside 'did you mean' suggestions.

For more information run: ```qcka --help```

*TODO: [generate markdown documentation using cobra](https://github.com/spf13/cobra/blob/main/site/content/docgen/md.md).*

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [ ] Alias valiadation
- [ ] Documentation
- [ ] Test suite
- [ ] Build pipeline
- [ ] Package & release

See the [open issues](https://github.com/SamTheSomebody/qcka/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Top contributors:

<a href="https://github.com/SamTheSomebody/qcka/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=samthesomebody/qcka" alt="contrib.rocks image" />
</a>

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GNU GPLv3 License.

See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact
@SamTheSomebody aka @GameDevSam

Sam Muller - gamedevsam@pm.me

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Cobra](https://github.com/spf13/cobra) for making cli tools so easy.
* othneildrew for the [best README template](https://github.com/othneildrew/Best-README-Template).
* [Boot.dev](https://www.boot.dev) for giving me the skills to make this.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/samthesomebody/qcka.svg?style=for-the-badge
[contributors-url]: https://github.com/SamTheSomebody/qcka/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/samthesomebody/qcka.svg?style=for-the-badge
[forks-url]: https://github.com/SamTheSomebody/qcka/network/members
[stars-shield]: https://img.shields.io/github/stars/samthesomebody/qcka.svg?style=for-the-badge
[stars-url]: https://github.com/SamTheSomebody/qcka/stargazers
[issues-shield]: https://img.shields.io/github/issues/samthesomebody/qcka.svg?style=for-the-badge
[issues-url]: https://github.com/SamTheSomebody/qcka/issues
[license-shield]: https://img.shields.io/github/license/samthesomebody/qcka.svg?style=for-the-badge
[license-url]: https://github.com/SamTheSomebody/qcka/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/gamedevsam/