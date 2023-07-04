# Patches I've Created so far.

## Fan control

### Labeled x-axis
Adds ticks to the x-axis that includes the tempature.
#### Known issues 
 - Causes the overall size of the chart to shrink slightly.
 - Looks a bit odd

Edit `C:\Program Files\Intel\Intel Arc Control\resource\js\chart_configs\charts.js`and replace
```JS
            x: {
                display: false,
                grid: {
                    display: false,
                },
                ticks: {
                    color: energyBlue,
                },
            },
```
with
```JS
            x: {
                display: true,
                grid: {
                    display: false,
                },
                ticks: {
                    color: energyBlue,
                    callback: function(value, index, values) {
                        return (value * 15) + 25 + "°C";
                    },
                    padding: 0, 
                    font: {
                        size: 15,
                        family: "IntelOneDisplay-en",
                    },
                    color: "#FFFFFF",
                },
                scaleLabel: {
                    display: true,
                },
            },
```
And under `C:\Program Files\Intel\Intel Arc Control\resource\js\pages\performance\performance_tuning.js` delete
```JS
document.getElementById('fan-graph-x-max').innerHTML = 100 + getTranslationFromId('units-celsius');
document.getElementById('fan-graph-x-min').innerHTML = 25 + getTranslationFromId('units-celsius');
```
## Notifications

### Disable driver update check timeout
Edit `C:\Program Files\Intel\Intel Arc Control\resource\js\pages\drivers\updates.js`
Delete this, but keep the `checkingForUpdates = false;` part.
```JS
        // dsa still hasn't returned to the frontend, cancel the call and let the user know.
        checkingForUpdates = false;
        showToast({
            type: notificationTypes.error,
            toggleType: notificationToggleTypes.notification_driver_info,
            mainMessageId: 'main-drivers',
            secondaryMessageId: 'drivers-checking-for-updates-timeout',
        });
```

## Minimal overlay
Edit `C:\Program Files\Intel\Intel Arc Control\resource\js\overlay.js`
Change
```JS
            <li id="${setting?.settingId}-wrapper" class="${hidden ? 'is-hidden' : ''}">
                <div class="left translatable" id="${setting?.settingId}">${setting?.name}</div>
                <div class="right unit-wrapper ${setting?.units == "units-percent" ? "percent" : ""}">
                    <span class="value" id="${setting?.settingId}-value">-</span>
                    <span class="unit translatable" id="${setting?.units}">FPS</span>
                </div>
            </li>`;
```
to
```JS
            <li id="${setting?.settingId}-wrapper" class="${hidden ? 'is-hidden' : ''}" style="margin: 0; padding: 0.25em;">
                <div class="left translatable" id="${setting?.settingId}">${setting?.name}</div>
                <div class="right unit-wrapper ${setting?.units == "units-percent" ? "percent" : ""}">
                    <span class="value" id="${setting?.settingId}-value">-</span>
                    <span class="unit translatable" id="${setting?.units}">FPS</span>
                </div>
            </li>`;
```