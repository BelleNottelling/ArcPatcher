# Patches I've Created so far.

## Performance tuning

### Better fan control

Adds ticks to the x-axis that includes the tempature, increases the size, and adds additional curve points.
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
          callback: function (value, index, values) {
            if (!Number.isInteger(value)) {
              return "";
            }
            length = isEmpty(activeOverclockingSettings()?.fan_speed_table) ? 12 : activeOverclockingSettings()?.fan_speed_table?.length;
            interval = 75 / (length - 1);
            return Math.round((value * interval + 25) * 2) / 2 + "°C";
          },
          padding: 0,
          font: {
            size: 14,
          },
          color: "#FFFFFF",
        },
        scaleLabel: {
          display: true,
        },
      },
```
And under `C:\Program Files\Intel\Intel Arc Control\resource\js\pages\performance\performance_tuning.js`
delete
```JS
document.getElementById('fan-graph-x-max').innerHTML = 100 + getTranslationFromId('units-celsius');
document.getElementById('fan-graph-x-min').innerHTML = 25 + getTranslationFromId('units-celsius');
```
repalce 
```JS
if (isEmpty(activeOverclockingSettings()?.fan_speed_table)) {
    activeOverclockingSettings().fan_speed_table = [30, 30, 40, 55, 75, 90];
}
```
with
```JS
if (isEmpty(activeOverclockingSettings()?.fan_speed_table) || activeOverclockingSettings()?.fan_speed_table?.length <= 6) {
    activeOverclockingSettings().fan_speed_table = [30, 30, 30, 30, 30, 55, 65, 75, 82.5, 90]; //[30, 30, 40, 55, 75, 90]
}
```

Finally edit `C:\Program Files\Intel\Intel Arc Control\resource\index.html`
Replace
```HTML
<div id="fan-speed-graph-container">
    <div id="fan-speed-graph-blockout"></div>
    <canvas id="fan-speed-graph-dragable"
    height="100"
    width="527"
    style="display: block; box-sizing: border-box; height: 100px; width: 527px; touch-action: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0); transform: translateZ(10px);"></canvas>
</div>
```
With
```HTML
<div id="fan-speed-graph-container">
    <canvas id="fan-speed-graph-dragable"
    height="125"
    width="527"
    style="display: block; box-sizing: border-box; height: 125px; width: 527px; touch-action: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0); transform: translateZ(10px);"></canvas>
</div>
```

### Performance Boost Show Mhz

Edit `C:\Program Files\Intel\Intel Arc Control\resource\js\pages\performance\performance_tuning.js`
Replace
```JS
visible: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.bSupported,
        
range: {
    normalize: true,
    min: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.min,
    max: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.max,
    step: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.step,
    defaultValue: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.default,
    units: null,
}
```
with
```JS
visible: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.bSupported,
showPosNeg: true,
        
range: {
    normalize: false,
    min: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.min,
    max: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.max,
    step: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.step,
    defaultValue: activePerformanceAdapter()?.supported_oc_features?.gpu_performance_boost?.default,
    units: 'units-mhz',
}
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

## Overlay

## Minimal Overaly
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

Edit `C:\Program Files\Intel\Intel Arc Control\resource\overlay.html`
Delete
```HTML
<div class="header">
    <div id="intel-performance-telemetry" class="left translatable">Intel&reg; Performance Telemetry</div>
</div>
```