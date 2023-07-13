# Archived 7-13-2023
After 4 months of dealing with lots of small and annoying issues which really shouldn't exist & being disappointed in Intel's support, I am selling my A770 and I do not wish to own another Intel GPU anytime soon.
This patcher will be archived and will not be getting any updates from me, although other people are welcome to make a fork of it and continue to update it.

# ArcPatcher
A handy utility to patch Intel's "Arc Control" to make it better.

## Available Patches
 - Better Fan Control
	 - Increases the number of points in the fan curve, makes the fan curve editor slightly larger, and also updates it to display the temperature at each point rather than only 25c and 100c on the left and right side.
     - *Note*: In order to increase the number of points in the fan curve, this patch will cause you to loose your existing fan curve.
	 - [Before](https://github.com/BelleNottelling/ArcPatcher/blob/main/Screenshots/originalFanControl.png) VS [after](https://github.com/BelleNottelling/ArcPatcher/blob/main/Screenshots/betterFanControl.png) 
- Minimized Overlay
	- Decreases the padding between overlay elements and removes the "Intel" branding from it.
	- [Before](https://github.com/BelleNottelling/ArcPatcher/blob/main/Screenshots/originalOverlay.png) VS [after](https://github.com/BelleNottelling/ArcPatcher/blob/main/Screenshots/minimalOverlay.png)
- Driver update timeout notification removal
- Performance boost display MHz
    - This patch changes the performance boost slider to display the actual MHz offset you are applying rather than displaying it as a percentage. (Default is a percentage of the possible MHz offset, which is typically 300Mhz for the A770 / A750)
- Arc Control Bug Fixes
	- A generic patch that will encopmas any Arc Control bugs I've created fixes for.
	- Fixes `applyPerformanceTuning does not update performance sliders`


## Usage
1. Exit Arc Control entirely. (`Task bar` -> `hidden icons` -> right click `Arc Control` -> `Quit Arc Control`)
2. Download the [latest release](https://github.com/BelleNottelling/ArcPatcher/releases), extract the archive, and then run `ArchPatcher.exe` as administrator.
	a. ArcPatcher will automatically create a backup of the directory that it will modify. 
3. From there you will be presented with a list of available patches, simply enter the number associated with the patch you'd like to apply.
4. Re-launch Arc Control and enjoy the patches!

## License
ArcPatcher is licensed under the [Apache 2.0](https://github.com/BelleNottelling/ArcPatcher/blob/main/LICENSE) license and comes under no warrenty.
