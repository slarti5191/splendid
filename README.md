As you will no doubt be aware, the plans for development of this project require a complete rewrite.  And regrettably you have arrived at the wrong repository.  This repository is one of those scheduled for demolition. The process will take approximately one of your cyber clicks right [here](https://github.com/aja-video/contra)
=====================================================
Splendid has been scrapped in favor of a complete rewrite, now known as Contra.

This repo is not in fact scheduled for demolition, however future updates are unlikely.

Check out Contra [here](https://github.com/aja-video/contra)

I don't know, apathetic bloody project, I've no sympathy at all.

Splendid - Configuration Tracking for Network Devices
=====================================================

Splendid helps track changes and backup network device configurations.

[![Build Status](https://travis-ci.org/slarti5191/splendid.png)](https://travis-ci.org/slarti5191/splendid) [![Go Report Card](https://goreportcard.com/badge/github.com/slarti5191/splendid)](https://goreportcard.com/report/github.com/slarti5191/splendid)

Features
--------

- Automated config watching to all supported network devices.
- Track changes in GIT.
- Review current status via web portal.

Devices Supported
-----------------

- Cisco CSB
- PFSense

ToDo:

- Cisco (Older models.)
- Comware
- External
- Juniper (JunOS)
- Vyatta

Device Configuration
--------------------
If a password, or any field in the config file, contains a `#` or `;` character be sure to properly
quote the password with either a backtick ``` ` ``` or a set of three double-quotes ``` """ ``` for
example, if your password is `Some#pass;word` you will need one of the following formats:

    Pass=`Some#pass;word`
    Pass="""Some#pass;word"""

License
-------

Splendid is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

Acknowledgments
---------------

- This project was originally inspired by Sweet and Rancid.
